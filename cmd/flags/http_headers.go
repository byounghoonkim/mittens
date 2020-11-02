//Copyright 2020 Expedia, Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package flags

import (
	"flag"
	"fmt"
)

// HTTPHeaders stores flags related to HTTP headers.
type HTTPHeaders struct {
	Headers stringArray
}

func (h *HTTPHeaders) String() string {
	return fmt.Sprintf("%+v", *h)
}

func (h *HTTPHeaders) initFlags() {
	flag.Var(&h.Headers, "http-headers", "HTTP header to be sent with warm up requests.")
}

func (h *HTTPHeaders) getWarmupHTTPHeaders() []string {
	return h.Headers
}
