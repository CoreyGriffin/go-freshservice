package freshservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const agentURL = "/api/v2/agents"

// AgentService is an interface for interacting with
// the agent endpoints of the Freshservice API
type AgentService interface {
	List(context.Context, QueryFilter) ([]AgentDetails, string, error)
	Create(context.Context, *AgentDetails) (*AgentDetails, error)
	Get(context.Context, int) (*AgentDetails, error)
	Update(context.Context, int, *AgentDetails) (*AgentDetails, error)
	Delete(context.Context, int) error
	Deactivate(context.Context, int) (*AgentDetails, error)
	Reactivate(context.Context, int) (*AgentDetails, error)
	ConvertToRequester(context.Context, int) (*AgentDetails, error)
}

// AgentServiceClient facilitates requests with the AgentService methods
type AgentServiceClient struct {
	client *Client
}

// List all freshservice agents
func (as *AgentServiceClient) List(ctx context.Context, filter QueryFilter) ([]AgentDetails, string, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   agentURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Agents{}
	resp, err := as.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}

// Get a specific Freshservice agent
func (as *AgentServiceClient) Get(ctx context.Context, id int) (*AgentDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/%d", agentURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Agent{}
	if _, err := as.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// Create a new Freshserrvice agent
func (as *AgentServiceClient) Create(ctx context.Context, ad *AgentDetails) (*AgentDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   agentURL,
	}

	agentContent, err := json.Marshal(ad)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(agentContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Agent{}
	if _, err := as.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// Update a Freshservice agent
func (as *AgentServiceClient) Update(ctx context.Context, id int, ad *AgentDetails) (*AgentDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/%d", agentURL, id),
	}

	agentContent, err := json.Marshal(ad)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(agentContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Agent{}
	if _, err := as.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil

}

// Delete a Freshservice agent
func (as *AgentServiceClient) Delete(ctx context.Context, id int) error {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/%d/forget", agentURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), nil)
	if err != nil {
		return err
	}

	if _, err := as.client.makeRequest(req, nil); err != nil {
		return err
	}

	return nil
}

// Deactivate a Frehservice agent (does not delete)
func (as *AgentServiceClient) Deactivate(ctx context.Context, id int) (*AgentDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/%d", agentURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Agent{}
	if _, err := as.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// Reactivate a Freshserrvice agent
func (as *AgentServiceClient) Reactivate(ctx context.Context, id int) (*AgentDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/%d", agentURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Agent{}
	if _, err := as.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// ConvertToRequester will convert a Freshservice agent to a requester
func (as *AgentServiceClient) ConvertToRequester(ctx context.Context, id int) (*AgentDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/%d/convert_to_requester", agentURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Agent{}
	if _, err := as.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
