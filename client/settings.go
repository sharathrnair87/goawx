package awx

import (
	"bytes"
	"encoding/json"
)

type SettingsService struct {
	client *Client
}

// Azure AD Oauth 2
const azureOauth2APIEndpoint = "/api/v2/settings/azuread-oauth2/"

func (s *SettingsService) CreateSettingsAzureADOauth2(data map[string]interface{}, params map[string]string) (*SettingsAzureADOauth2, error) {
	result := new(SettingsAzureADOauth2)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PutJSON(azureOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) GetSettingsAzureADOauth2(params map[string]string) (*SettingsAzureADOauth2, error) {
	result := new(SettingsAzureADOauth2)
	resp, err := s.client.Requester.GetJSON(azureOauth2APIEndpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) UpdateSettingsAzureADOauth2(data map[string]interface{}, params map[string]string) (*SettingsAzureADOauth2, error) {
	result := new(SettingsAzureADOauth2)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PatchJSON(azureOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) DeleteSettingsAzureADOauth2(params map[string]string) error {
	resp, err := s.client.Requester.Delete(azureOauth2APIEndpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}

// Github.com Oauth2
const githubOauth2APIEndpoint = "/api/v2/settings/github/"

func (s *SettingsService) CreateSettingsGithubOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubOauth2, error) {
	result := new(SettingsGithubOauth2)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PutJSON(githubOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) GetSettingsGithubOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubOauth2, error) {
	result := new(SettingsGithubOauth2)
	resp, err := s.client.Requester.GetJSON(githubOauth2APIEndpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) UpdateSettingsGithubOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubOauth2, error) {
	result := new(SettingsGithubOauth2)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PatchJSON(githubOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) DeleteSettingsGithubOauth2(params map[string]string) error {
	resp, err := s.client.Requester.Delete(githubOauth2APIEndpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}

// GitHub Enterprise OAuth2
const githubEnterpriseOauth2APIEndpoint = "/api/v2/settings/github-enterprise/"

func (s *SettingsService) CreateSettingsGithubEnterpriseOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubEnterpriseOauth2, error) {
	result := new(SettingsGithubEnterpriseOauth2)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PutJSON(githubEnterpriseOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) GetSettingsGithubEnterpriseOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubEnterpriseOauth2, error) {
	result := new(SettingsGithubEnterpriseOauth2)
	resp, err := s.client.Requester.GetJSON(githubEnterpriseOauth2APIEndpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) UpdateSettingsGithubEnterpriseOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubEnterpriseOauth2, error) {
	result := new(SettingsGithubEnterpriseOauth2)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PatchJSON(githubEnterpriseOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) DeleteSettingsGithubEnterpriseOauth2(params map[string]string) error {
	resp, err := s.client.Requester.Delete(githubEnterpriseOauth2APIEndpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}

// GitHub Enterprise Organization OAuth2
const githubEnterpriseOrgOauth2APIEndpoint = "/api/v2/settings/github-enterprise-org/"

func (s *SettingsService) CreateSettingsGithubEnterpriseOrgOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubEnterpriseOrgOauth2, error) {
	result := new(SettingsGithubEnterpriseOrgOauth2)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PutJSON(githubEnterpriseOrgOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) GetSettingsGithubEnterpriseOrgOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubEnterpriseOrgOauth2, error) {
	result := new(SettingsGithubEnterpriseOrgOauth2)
	resp, err := s.client.Requester.GetJSON(githubEnterpriseOrgOauth2APIEndpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) UpdateSettingsGithubEnterpriseOrgOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubEnterpriseOrgOauth2, error) {
	result := new(SettingsGithubEnterpriseOrgOauth2)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PatchJSON(githubEnterpriseOrgOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) DeleteSettingsGithubEnterpriseOrgOauth2(params map[string]string) error {
	resp, err := s.client.Requester.Delete(githubEnterpriseOrgOauth2APIEndpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}

// GitHub Enterprise Team OAuth2
const githubEnterpriseTeamOauth2APIEndpoint = "/api/v2/settings/github-enterprise-team/"

func (s *SettingsService) CreateSettingsGithubEnterpriseTeamOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubEnterpriseTeamOauth2, error) {
	result := new(SettingsGithubEnterpriseTeamOauth2)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PutJSON(githubEnterpriseTeamOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) GetSettingsGithubEnterpriseTeamOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubEnterpriseTeamOauth2, error) {
	result := new(SettingsGithubEnterpriseTeamOauth2)
	resp, err := s.client.Requester.GetJSON(githubEnterpriseTeamOauth2APIEndpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) UpdateSettingsGithubEnterpriseTeamOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubEnterpriseTeamOauth2, error) {
	result := new(SettingsGithubEnterpriseTeamOauth2)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PatchJSON(githubEnterpriseTeamOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) DeleteSettingsGithubEnterpriseTeamOauth2(params map[string]string) error {
	resp, err := s.client.Requester.Delete(githubEnterpriseTeamOauth2APIEndpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}

// GitHub Organization OAuth2
const githubOrgOauth2APIEndpoint = "/api/v2/settings/github-org/"

func (s *SettingsService) CreateSettingsGithubOrgOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubOrgOauth2, error) {
	result := new(SettingsGithubOrgOauth2)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PutJSON(githubOrgOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) GetSettingsGithubOrgOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubOrgOauth2, error) {
	result := new(SettingsGithubOrgOauth2)
	resp, err := s.client.Requester.GetJSON(githubOrgOauth2APIEndpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) UpdateSettingsGithubOrgOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubOrgOauth2, error) {
	result := new(SettingsGithubOrgOauth2)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PatchJSON(githubOrgOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) DeleteSettingsGithubOrgOauth2(params map[string]string) error {
	resp, err := s.client.Requester.Delete(githubOrgOauth2APIEndpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}

// GitHub Team OAuth2
const githubTeamOauth2APIEndpoint = "/api/v2/settings/github-team/"

func (s *SettingsService) CreateSettingsGithubTeamOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubTeamOauth2, error) {
	result := new(SettingsGithubTeamOauth2)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PutJSON(githubTeamOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) GetSettingsGithubTeamOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubTeamOauth2, error) {
	result := new(SettingsGithubTeamOauth2)
	resp, err := s.client.Requester.GetJSON(githubTeamOauth2APIEndpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) UpdateSettingsGithubTeamOauth2(data map[string]interface{}, params map[string]string) (*SettingsGithubTeamOauth2, error) {
	result := new(SettingsGithubTeamOauth2)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PatchJSON(githubTeamOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) DeleteSettingsGithubTeamOauth2(params map[string]string) error {
	resp, err := s.client.Requester.Delete(githubTeamOauth2APIEndpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}

// Google OAuth2
const GoogleOauth2APIEndpoint = "/api/v2/settings/google-oauth2/"

func (s *SettingsService) CreateSettingsGoogleOauth2(data map[string]interface{}, params map[string]string) (*SettingsGoogleOauth2, error) {
	result := new(SettingsGoogleOauth2)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PutJSON(GoogleOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) GetSettingsGoogleOauth2(data map[string]interface{}, params map[string]string) (*SettingsGoogleOauth2, error) {
	result := new(SettingsGoogleOauth2)
	resp, err := s.client.Requester.GetJSON(GoogleOauth2APIEndpoint, result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) UpdateSettingsGoogleOauth2(data map[string]interface{}, params map[string]string) (*SettingsGoogleOauth2, error) {
	result := new(SettingsGoogleOauth2)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Requester.PatchJSON(GoogleOauth2APIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SettingsService) DeleteSettingsGoogleOauth2(params map[string]string) error {
	resp, err := s.client.Requester.Delete(GoogleOauth2APIEndpoint, nil, params)
	if err != nil {
		return err
	}

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}
