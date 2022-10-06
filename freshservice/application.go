package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const applicationURL = "/api/v2/applications"

// ApplicationService is an interface for interacting with
// the application endpoints of the Freshservice API
type ApplicationService interface {
	List(context.Context, QueryFilter) ([]ApplicationDetails, string, error)
	Get(context.Context, int64) (*ApplicationDetails, error)
	ListLicenses(context.Context, int64) ([]LicensesDetails, error)
	ListUsers(context.Context, int64) ([]ApplicationUserDetails, error)
	ListInstallations(context.Context, int64) ([]ApplicationInstallationDetails, error)
}

// ApplicationServiceClient facilitates requests with the TicketService methods
type ApplicationServiceClient struct {
	client *Client
}

// List all application
// All the below requests are paginated to return only 30 tickets per page.
// Append the parameter "page=[:page_no]" in the url to traverse through pages.
func (a *ApplicationServiceClient) List(ctx context.Context, filter QueryFilter) ([]ApplicationDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   applicationURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Applications{}
	resp, err := a.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}

// Get a specific all application
func (a *ApplicationServiceClient) Get(ctx context.Context, appID int64) (*ApplicationDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d", applicationURL, appID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Application{}
	if _, err = a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// ListLicenses lists all the licenses for an application
func (a *ApplicationServiceClient) ListLicenses(ctx context.Context, appID int64) ([]LicensesDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d/licenses", applicationURL, appID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Licenses{}
	if _, err = a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res.List, nil
}

// ListUsers lists all the users of an application
func (a *ApplicationServiceClient) ListUsers(ctx context.Context, appID int64) ([]ApplicationUserDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d/users", applicationURL, appID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &ApplicationUsers{}
	if _, err = a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res.List, nil
}

// ListInstallations lists all the installations of an application
func (a *ApplicationServiceClient) ListInstallations(ctx context.Context, appID int64) ([]ApplicationInstallationDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d/installations", applicationURL, appID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &ApplicationInstallations{}
	if _, err = a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res.List, nil
}

// QueryString allows us to pass TicketListOptions as a QueryFilter and
// will return a new endpoint URL with query parameters attached
func (opts *ApplicationListOptions) QueryString() string {
	var qs []string

	if opts.PageQuery != "" {
		qs = append(qs, opts.PageQuery)
	}

	return strings.Join(qs, "&")
}
