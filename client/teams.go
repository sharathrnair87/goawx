package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// TeamService implements awx teams apis.
type TeamService struct {
	client *Client
}

// ListTeamsResponse represents `ListTeams` endpoint response.
type ListTeamsResponse struct {
	Pagination
	Results []*Team `json:"results"`
}

type ListTeamRoleEntitlementsResponse struct {
	Pagination
	Results []*ApplyRole `json:"results"`
}

const teamsAPIEndpoint = "/api/v2/teams/"

// ListTeams shows list of awx teams.
func (p *TeamService) ListTeams(params map[string]string) ([]*Team, *ListTeamsResponse, error) {
	result := new(ListTeamsResponse)
	resp, err := p.client.Requester.GetJSON(teamsAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

func (p *TeamService) ListTeamRoleEntitlements(id int, params map[string]string) ([]*ApplyRole, *ListTeamRoleEntitlementsResponse, error) {
	result := new(ListTeamRoleEntitlementsResponse)
	endpoint := fmt.Sprintf("%s%d/roles/", teamsAPIEndpoint, id)
	resp, err := p.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}
	return result.Results, result, nil
}

// GetTeamByID shows the details of a team.
func (p *TeamService) GetTeamByID(id int, params map[string]string) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("%s%d/", teamsAPIEndpoint, id)
	resp, err := p.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateTeam creates an awx team.
func (p *TeamService) CreateTeam(data map[string]interface{}, params map[string]string) (*Team, error) {
	mandatoryFields = []string{"name", "organization"}
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Team)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Add check if team exists and return proper error

	resp, err := p.client.Requester.PostJSON(teamsAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateTeam update an awx Team.
func (p *TeamService) UpdateTeam(id int, data map[string]interface{}, params map[string]string) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("%s%d/", teamsAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

func (p *TeamService) UpdateTeamRoleEntitlement(id int, data map[string]interface{}, params map[string]string) (interface{}, error) {
	result := new(interface{})
	endpoint := fmt.Sprintf("%s%d/roles/", teamsAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteTeam delete an awx Team.
func (p *TeamService) DeleteTeam(id int) (*Team, error) {
	result := new(Team)
	endpoint := fmt.Sprintf("%s%d", teamsAPIEndpoint, id)

	resp, err := p.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
