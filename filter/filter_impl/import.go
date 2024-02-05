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

// Package filter_impl is for being compatible with older dubbo-go, please use `imports` package.
// It may be DEPRECATED OR REMOVED in the future.
package filter_impl

import (
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
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/seata"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/sentinel"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/token"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/tps"
	_ "github.com/xd-luqiang/dubbo-go/v3/filter/tracing"
)
