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
