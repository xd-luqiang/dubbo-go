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

package failfast

import (
	clusterpkg "github.com/xd-luqiang/dubbo-go/v3/cluster/cluster"
	"github.com/xd-luqiang/dubbo-go/v3/cluster/directory"
	"github.com/xd-luqiang/dubbo-go/v3/common/constant"
	"github.com/xd-luqiang/dubbo-go/v3/common/extension"
	"github.com/xd-luqiang/dubbo-go/v3/protocol"
)

func init() {
	extension.SetCluster(constant.ClusterKeyFailfast, newFailfastCluster)
}

type failfastCluster struct{}

// newFailfastCluster returns a failfastCluster instance.
//
// Fast failure, only made a call, failure immediately error. Usually used for non-idempotent write operations,
// such as adding records.
func newFailfastCluster() clusterpkg.Cluster {
	return &failfastCluster{}
}

// Join returns a baseClusterInvoker instance
func (cluster *failfastCluster) Join(directory directory.Directory) protocol.Invoker {
	return clusterpkg.BuildInterceptorChain(newFailfastClusterInvoker(directory))
}
