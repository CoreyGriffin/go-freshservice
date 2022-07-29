package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const businessHoursURL = "/api/v2/business_hours"

// BusinessHoursService is an interface for interacting with
// the business hours endpoints of the Freshservice API
type BusinessHoursService interface {
	List(context.Context) ([]BusinessHoursDetails, error)
	Get(context.Context, int) (*BusinessHoursDetails, error)
}

// BusinessHoursServiceClient facilitates requests with the AnnouncementService methods
type BusinessHoursServiceClient struct {
	client *Client
}

// List all business hours configured in Freshservice
func (c *BusinessHoursServiceClient) List(ctx context.Context) ([]BusinessHoursDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.Domain,
		Path:   businessHoursURL,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &BusinessHours{}
	if _, err := c.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res.List, nil
}

// Get a details for a specific business hour configuration in Freshservice
func (c *BusinessHoursServiceClient) Get(ctx context.Context, id int) (*BusinessHoursDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.Domain,
		Path:   fmt.Sprintf("%s/%d", businessHoursURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &BusinessHoursConfig{}
	if _, err := c.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}
