package awx

type CredentialTypeService struct {
	client *Client
}

type ListCredentialTypeResponse struct {
	Pagination
	Results []*CredentialType `json:"results"`
}

func (cs *CredentialTypeService) ListCredentialsTypes(params map[string]string) ([]*CredentialType,
	*ListCredentialTypeResponse, error) {
	result := new(ListCredentialTypeResponse)
	endpoint := "/api/v2/credential_types/"
	resp, err := cs.client.Requester.GetJSON(endpoint, result, params)
	if err != nil { return nil, result, err}

	err = CheckResponse(resp)
	if err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}