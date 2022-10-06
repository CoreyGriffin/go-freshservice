package freshservice

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const assetURL = "/api/v2/assets"

// AssetService is an interface for interacting with
// the asset endpoints of the Freshservice API
type AssetService interface {
	List(context.Context, QueryFilter) ([]AssetDetails, string, error)
	Get(context.Context, int) (*AssetDetails, error)
}

// AssetServiceClient facilitates requests with the AssetService methods
type AssetServiceClient struct {
	client *Client
}

// List all Assets
// Append the parameter "page=[:page_no]" in the url to traverse through pages.
func (a *AssetServiceClient) List(ctx context.Context, filter QueryFilter) ([]AssetDetails, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   assetURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Assets{}
	resp, err := a.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}

// Get a specific asset
func (a *AssetServiceClient) Get(ctx context.Context, assetID int) (*AssetDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   a.client.Domain,
		Path:   fmt.Sprintf("%s/%d", assetURL, assetID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Asset{}
	if _, err = a.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return &res.Details, nil
}

// QueryString allows us to pass AssetListOptions as a QueryFilter and
// will return a new endpoint URL with query parameters attached
func (opts *AssetListOptions) QueryString() string {
	var qs []string

	if opts.PageQuery != "" {
		qs = append(qs, opts.PageQuery)
	}

	if opts.Embed != nil {
		if opts.Embed.TypeFields {
			qs = append(qs, "include=type_fields")
		}
		if opts.Embed.Trashed {
			qs = append(qs, fmt.Sprintf("trashed=%v", opts.Embed.Trashed))
		}
	}

	return strings.Join(qs, "&")
}
