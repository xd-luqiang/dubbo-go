/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 *
 * Copyright 2020 gRPC authors.
 *
 */

// Package v3 provides xDS v3 transport protocol specific functionality.
package v3

import (
	"context"
	"fmt"
)

import (
	v3corepb "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3adsgrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	v3discoverypb "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"

	"github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"google.golang.org/protobuf/types/known/anypb"
)

import (
	dubboLogger "github.com/xd-luqiang/dubbo-go/v3/common/logger"
	controllerversion "github.com/xd-luqiang/dubbo-go/v3/xds/client/controller/version"
	"github.com/xd-luqiang/dubbo-go/v3/xds/client/resource"
	resourceversion "github.com/xd-luqiang/dubbo-go/v3/xds/client/resource/version"
	"github.com/xd-luqiang/dubbo-go/v3/xds/utils/pretty"
)

func init() {
	controllerversion.RegisterAPIClientBuilder(resourceversion.TransportV3, newClient)
}

var (
	resourceTypeToURL = map[resource.ResourceType]string{
		resource.ListenerResource:    resourceversion.V3ListenerURL,
		resource.RouteConfigResource: resourceversion.V3RouteConfigURL,
		resource.ClusterResource:     resourceversion.V3ClusterURL,
		resource.EndpointsResource:   resourceversion.V3EndpointsURL,
	}
)

func newClient(opts controllerversion.BuildOptions) (controllerversion.MetadataWrappedVersionClient, error) {
	nodeProto, ok := opts.NodeProto.(*v3corepb.Node)
	if !ok {
		return nil, fmt.Errorf("xds: unsupported Node proto type: %T, want %T", opts.NodeProto, v3corepb.Node{})
	}
	v3c := &client{
		nodeProto: nodeProto, logger: dubboLogger.GetLogger(),
	}
	return v3c, nil
}

type adsStream v3adsgrpc.AggregatedDiscoveryService_StreamAggregatedResourcesClient

// client performs the actual xDS RPCs using the xDS v3 API. It creates a
// single ADS stream on which the different types of xDS requests and responses
// are multiplexed.
type client struct {
	nodeProto *v3corepb.Node
	logger    dubboLogger.Logger
}

func (v3c *client) NewStream(ctx context.Context, cc *grpc.ClientConn) (grpc.ClientStream, error) {
	return v3adsgrpc.NewAggregatedDiscoveryServiceClient(cc).StreamAggregatedResources(ctx, grpc.WaitForReady(true))
}

// SetMetadata update client metadata
func (v3c *client) SetMetadata(p *_struct.Struct) {
	v3c.nodeProto.Metadata = p
}

// SendRequest sends out a DiscoveryRequest for the given resourceNames, of type
// rType, on the provided stream.
//
// version is the ack version to be sent with the request
// - If this is the new request (not an ack/nack), version will be empty.
// - If this is an ack, version will be the version from the response.
// - If this is a nack, version will be the previous acked version (from
//   versionMap). If there was no ack before, it will be empty.
func (v3c *client) SendRequest(s grpc.ClientStream, resourceNames []string, rType resource.ResourceType, version, nonce, errMsg string) error {
	stream, ok := s.(adsStream)
	if !ok {
		return fmt.Errorf("xds: Attempt to send request on unsupported stream type: %T", s)
	}
	req := &v3discoverypb.DiscoveryRequest{
		Node:          v3c.nodeProto,
		TypeUrl:       resourceTypeToURL[rType],
		ResourceNames: resourceNames,
		VersionInfo:   version,
		ResponseNonce: nonce,
	}
	if errMsg != "" {
		req.ErrorDetail = &statuspb.Status{
			Code: int32(codes.InvalidArgument), Message: errMsg,
		}
	}
	if err := stream.Send(req); err != nil {
		return fmt.Errorf("xds: stream.Send(%+v) failed: %v", req, err)
	}
	v3c.logger.Debugf("ADS request sent: %v", pretty.ToJSON(req))
	return nil
}

// RecvResponse blocks on the receipt of one response message on the provided
// stream.
func (v3c *client) RecvResponse(s grpc.ClientStream) (proto.Message, error) {
	stream, ok := s.(adsStream)
	if !ok {
		return nil, fmt.Errorf("xds: Attempt to receive response on unsupported stream type: %T", s)
	}

	resp, err := stream.Recv()
	if err != nil {
		return nil, fmt.Errorf("xds: stream.Recv() failed: %v", err)
	}
	v3c.logger.Infof("ADS response received, type: %v", resp.GetTypeUrl())
	v3c.logger.Debugf("ADS response received: %+v", pretty.ToJSON(resp))
	return resp, nil
}

func (v3c *client) ParseResponse(r proto.Message) (resource.ResourceType, []*anypb.Any, string, string, error) {
	rType := resource.UnknownResource
	resp, ok := r.(*v3discoverypb.DiscoveryResponse)
	if !ok {
		return rType, nil, "", "", fmt.Errorf("xds: unsupported message type: %T", resp)
	}

	// Note that the xDS transport protocol is versioned independently of
	// the resource types, and it is supported to transfer older versions
	// of resource types using new versions of the transport protocol, or
	// vice-versa. Hence we need to handle v3 type_urls as well here.
	var err error
	url := resp.GetTypeUrl()
	switch {
	case resource.IsListenerResource(url):
		rType = resource.ListenerResource
	case resource.IsRouteConfigResource(url):
		rType = resource.RouteConfigResource
	case resource.IsClusterResource(url):
		rType = resource.ClusterResource
	case resource.IsEndpointsResource(url):
		rType = resource.EndpointsResource
	default:
		return rType, nil, "", "", controllerversion.ErrResourceTypeUnsupported{
			ErrStr: fmt.Sprintf("Resource type %v unknown in response from server", resp.GetTypeUrl()),
		}
	}
	return rType, resp.GetResources(), resp.GetVersionInfo(), resp.GetNonce(), err
}
