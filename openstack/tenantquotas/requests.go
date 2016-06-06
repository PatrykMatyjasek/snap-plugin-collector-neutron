/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

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
package tenantquotas

import (
	"net/http"

	"github.com/rackspace/gophercloud"
)

const (
	quotasPath = "quotas"
)

// Get will retrieve the volume type with the provided ID. To extract the volume
// type from the result, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, tenant string) Result {
	var res Result
	reqOpts := gophercloud.RequestOpts{
		OkCodes: []int{http.StatusOK, http.StatusMultipleChoices},
	}
	url := client.ServiceURL(quotasPath, tenant)
	_, res.Err = client.Get(url, &res.Body, &reqOpts)
	return res
}
