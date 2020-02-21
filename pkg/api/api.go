/*
Copyright 2020, Staffbase GmbH and contributors.

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

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const endpoint = "https://api.cloud.syseleven.net:5001"

func GetQuota(projectID, token string) (map[string]Quota, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/projects/%s/quota", endpoint, projectID), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Auth-Token", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var quotas map[string]Quota
	quotas = make(map[string]Quota)

	err = json.NewDecoder(resp.Body).Decode(&quotas)
	if err != nil {
		return nil, err
	}

	return quotas, nil
}

func GetCurrentUsage(projectID, token string) (map[string]CurrentUsage, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/projects/%s/current_usage", endpoint, projectID), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Auth-Token", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var currentUsages map[string]CurrentUsage
	currentUsages = make(map[string]CurrentUsage)

	err = json.NewDecoder(resp.Body).Decode(&currentUsages)
	if err != nil {
		return nil, err
	}

	return currentUsages, nil
}
