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

package imports

import (
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/cluster/adaptivesvc"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/cluster/available"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/cluster/broadcast"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/cluster/failback"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/cluster/failfast"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/cluster/failover"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/cluster/failsafe"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/cluster/forking"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/cluster/zoneaware"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/loadbalance/consistenthashing"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/loadbalance/leastactive"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/loadbalance/p2c"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/loadbalance/random"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/loadbalance/roundrobin"
	_ "github.com/xd-luqiang/dubbo-go/v3/cluster/router/meshrouter"
	_ "github.com/xd-luqiang/dubbo-go/v3/common/proxy/proxy_factory"
	_ "github.com/xd-luqiang/dubbo-go/v3/config_center/apollo"
	_ "github.com/xd-luqiang/dubbo-go/v3/config_center/nacos"
	_ "github.com/xd-luqiang/dubbo-go/v3/config_center/zookeeper"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/accesslog"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/active"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/adaptivesvc"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/auth"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/echo"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/exec_limit"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/generic"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/graceful_shutdown"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/hystrix"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/metrics"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/otel/trace"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/seata"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/sentinel"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/token"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/tps"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/tps/limiter"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/tps/strategy"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/tracing"
	_ "github.com/xd-luqiang/dubbo-go/v3/metadata/mapping/metadata"
	_ "github.com/xd-luqiang/dubbo-go/v3/metadata/report/etcd"
	_ "github.com/xd-luqiang/dubbo-go/v3/metadata/report/nacos"
	_ "github.com/xd-luqiang/dubbo-go/v3/metadata/report/zookeeper"
	_ "github.com/xd-luqiang/dubbo-go/v3/metadata/service/exporter/configurable"
	_ "github.com/xd-luqiang/dubbo-go/v3/metadata/service/local"
	_ "github.com/xd-luqiang/dubbo-go/v3/metadata/service/remote"
	_ "github.com/xd-luqiang/dubbo-go/v3/metrics/prometheus"
	_ "github.com/xd-luqiang/dubbo-go/v3/protocol/dubbo"
	_ "github.com/xd-luqiang/dubbo-go/v3/protocol/dubbo3"
	_ "github.com/xd-luqiang/dubbo-go/v3/protocol/dubbo3/reflection"
	_ "github.com/xd-luqiang/dubbo-go/v3/protocol/grpc"
	_ "github.com/xd-luqiang/dubbo-go/v3/protocol/jsonrpc"
	_ "github.com/xd-luqiang/dubbo-go/v3/protocol/rest"
	_ "github.com/xd-luqiang/dubbo-go/v3/registry/etcdv3"
	_ "github.com/xd-luqiang/dubbo-go/v3/registry/nacos"
	_ "github.com/xd-luqiang/dubbo-go/v3/registry/polaris"
	_ "github.com/xd-luqiang/dubbo-go/v3/registry/protocol"
	_ "github.com/xd-luqiang/dubbo-go/v3/registry/servicediscovery"
	_ "github.com/xd-luqiang/dubbo-go/v3/registry/xds"
	_ "github.com/xd-luqiang/dubbo-go/v3/registry/zookeeper"
	_ "github.com/xd-luqiang/dubbo-go/v3/xds/client/controller/version/v2"
	_ "github.com/xd-luqiang/dubbo-go/v3/xds/client/controller/version/v3"
)
