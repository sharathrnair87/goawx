package awx

import (
	"bytes"
	"encoding/json"
)

type SettingGithubEnterpriseService struct {
	client *Client
}

const settingsGithubEnterpriseAPIEndpoint = "/api/v2/settings/github-enterprise/"

func (setting *SettingGithubEnterpriseService) GetSettingsGithubEnterprise(params map[string]string) (*SettingsGithubEnterprise, error) {
	result := new(SettingsGithubEnterprise)
	endpoint := settingsGithubEnterpriseAPIEndpoint
	resp, err := setting.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (setting *SettingGithubEnterpriseService) UpdateSettingsGithubEnterprise(data map[string]interface{}, params map[string]string) (*SettingsGithubEnterprise, error) {
	result := new(SettingsGithubEnterprise)
	endpoint := settingsGithubEnterpriseAPIEndpoint

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := setting.client.Requester.PutJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (setting *SettingGithubEnterpriseService) DeleteSettingsGithubEnterprise(params map[string]string) error {
	endpoint := settingsGithubEnterpriseAPIEndpoint
	resp, err := setting.client.Requester.Delete(endpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}
