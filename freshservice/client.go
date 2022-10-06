package freshservice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// Client represents a new Freshservice API client to
// be utilized for API requests
type Client struct {
	// Freshservice domain
	Domain string
	// Context to leverage during the lifetime of the client
	Context context.Context

	// Basic Authentication requried for Freshservice API calls
	Auth *BasicAuth
	// API client to utilize for making HTTP requests
	client *http.Client
}

// BasicAuth holds the basic auth requirements needed to
// utilize the Freshservice API
type BasicAuth struct {
	APIKey string
}

// Used if custom client not passed in when NewClient instantiated
func defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Minute,
	}
}

// New returns a new Freshservice API client that can be used for both V1 and V2 of the Freshservice API
func New(ctx context.Context, domain string, apikey string, client *http.Client) (*Client, error) {

	if ctx == nil {
		ctx = context.Background()
	}

	// handle required attributes
	if domain == "" {
		return nil, missingClientConfigErr("domain")
	}

	if apikey == "" {
		return nil, missingClientConfigErr("API key")
	}

	// default to HTTP client if one is not provided
	if client == nil {
		client = defaultHTTPClient()
		client.Timeout = time.Minute * 5
	}

	return &Client{
		Domain:  stripURLScheme(domain),
		Context: ctx,
		Auth: &BasicAuth{
			APIKey: apikey,
		},
		client: client,
	}, nil
}

// makeRequest is used internally by the Freshservice API client to
// make an API request and unmarshal into the response interface passed in
func (fs *Client) makeRequest(r *http.Request, v interface{}) (*http.Response, error) {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0, post-check=0, pre-check=0")
	r.Header.Set("Strict-Transport-Security", "max-age=31536000 ; includeSubDomains")
	r.SetBasicAuth(fs.Auth.APIKey, "x")

	// Replace scheme for unit tests that are using a mock server
	if os.Getenv("GO_TEST") == "1" {
		r.URL.Scheme = "http"
	}

	r.Close = true

	res, err := fs.client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("error making %s request to %s", r.Method, r.URL)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if res.StatusCode == http.StatusNotFound {
		return res, fmt.Errorf("%s %s not found", r.Method, r.URL)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode > 299 {
		v = ErrorResponse{}
	}

	if v == nil {
		return res, nil
	}

	return res, json.NewDecoder(res.Body).Decode(&v)
}

// We set the scheme in the HTTP request
func stripURLScheme(domain string) string {
	domain = strings.Replace(domain, "https://", "", -1)
	domain = strings.Replace(domain, "http://", "", -1)
	return domain
}

// Application is the interface between the HTTP client and the Freshservice application related endpoints
func (fs *Client) Applications() ApplicationService {
	return &ApplicationServiceClient{client: fs}
}

// Asset is the interface between the HTTP client and the Freshservice asset related endpoints
func (fs *Client) Assets() AssetService {
	return &AssetServiceClient{client: fs}
}

// Tickets is the interface between the HTTP client and the Freshservice ticket related endpoints
func (fs *Client) Tickets() TicketService {
	return &TicketServiceClient{client: fs}
}

// ServiceCatalog is the interface between the HTTP client and the Freshservice service catalog related endpoints
func (fs *Client) ServiceCatalog() ServiceCatalogService {
	return &ServiceCatalogServiceClient{client: fs}
}

// Announcements is the interface between the HTTP client and the Freshservice announcement related endpoints
func (fs *Client) Announcements() AnnouncementService {
	return &AnnouncementServiceClient{client: fs}
}

// Agents is the interface between the HTTP client and the Freshservice agent related endpoints
func (fs *Client) Agents() AgentService {
	return &AgentServiceClient{client: fs}
}

// BusinessHours is the interface between the HTTP client and the Freshservice business hours related endpoints
func (fs *Client) BusinessHours() BusinessHoursService {
	return &BusinessHoursServiceClient{client: fs}
}

// Tasks is the interface between the HTTP client and the Freshservice business hours related endpoints
func (fs *Client) Tasks() TaskService {
	return &TaskServiceClient{client: fs}
}
