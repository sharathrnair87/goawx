package awx

import (
	"bytes"
	"encoding/json"
)

type CredentialsService struct {
	client *Client
}

type ListCredentialsResponse struct {
	Pagination
	Results []*Credential `json:"results"'`
}

type ListCredentialsTypeResponse struct {
	Pagination
	Results []*CredentialType `json:"results"`
}

func (cs *CredentialsService) ListCredentials(params map[string]string) ([]*Credential,
	*ListCredentialsResponse,
	error) {
	result := new(ListCredentialsResponse)
	endpoint := "/api/v2/credentials"
	resp, err := cs.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

func (cs *CredentialsService) ListCredentialsTypes(params map[string]string) ([]*CredentialType,
	*ListCredentialsTypeResponse, error) {
	result := new(ListCredentialsTypeResponse)
	endpoint := "/api/v2/credential_types/"
	resp, err := cs.client.Requester.GetJSON(endpoint, result, params)
	if err != nil { return nil, result, err}

	err = CheckResponse(resp)
	if err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

func (cs *CredentialsService) CreateCredentials(data map[string]interface{}, params map[string]string) (*Credential, error) {
	result := new(Credential)
	endpoint := "/api/v2/credentials/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := cs.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}
