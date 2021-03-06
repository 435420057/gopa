/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package http

import "github.com/medcl/gopa/core/api"

type API struct {
	api.Handler
}

func InitAPI() {

	apis := API{}
	api.HandleMethod(api.GET, "/joint/", apis.handleGetPipelineJointsRequest)
	api.HandleMethod(api.POST, "/joint/", apis.handlePostPipelineJointsRequest)
	api.HandleMethod(api.POST, "/pipeline/:id", apis.handlePostPipelinesRequest)
	api.HandleMethod(api.GET, "/pipeline/", apis.handleGetPipelinesRequest)
	api.HandleMethod(api.GET, "/user/", apis.handleUserLoginRequest)

}
