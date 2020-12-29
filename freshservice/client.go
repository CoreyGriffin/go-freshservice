package freshservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Client represents a new Freshservice API client to
// be utilized for API requests
type Client struct {
	// Freshservice domain
	Domain string
	// Version of the API to utilize (defaults to v1)
	Version int
	// endpoint is a calculated field based on the version number and is
	// either equal to the domain or domain/api/v2
	Endpoint string
	// Context to leverage during the lifetime of the client
	Context context.Context
	// Logging configuration
	Logger *logrus.Logger
	// Basic Authentication requried for Freshservice API calls
	Auth *BasicAuth
	// API client to utilize for making HTTP requests
	client *http.Client
}

// BasicAuth holds the basic auth requirements needed to
// utilize the Freshservice API
type BasicAuth struct {
	Username string
	APIKey   string
}

// Used if custom client not passed in when NewClient instantiated
func defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Minute,
	}
}

// New returns a new Freshservice API client that can be used for both V1 and V2 of the Freshservice API
func New(ctx context.Context, domain string, version int, username string, secret string, l *logrus.Logger, client *http.Client) (*Client, error) {

	if ctx == nil {
		ctx = context.Background()
	}

	// handle required attributes
	if domain == "" {
		return nil, missingClientConfigErr("domain")
	}

	if username == "" {
		return nil, missingClientConfigErr("username")
	}

	if secret == "" {
		return nil, missingClientConfigErr("API Key")
	}

	// default to V1 if an API version is not provided
	// and error out if version greater than 2 is provided
	var ep string
	switch version {
	case 1:
		ep = domain
	case 2:
		ep = fmt.Sprintf("%s/api/v2", domain)
	default:
		version = 1
		ep = domain
	}

	// default to HTTP client if one is not provided
	if client == nil {
		client = defaultHTTPClient()
		client.Timeout = time.Minute * 5
	}

	return &Client{
		Domain:   domain,
		Version:  version,
		Endpoint: ep,
		Context:  ctx,
		Logger:   l,
		Auth: &BasicAuth{
			Username: username,
			APIKey:   secret,
		},
		client: client,
	}, nil
}

// makeRequest is used internally by the Freshservice API client to
// make an API request and unmarshal into the response interface passed in
func (fs *Client) makeRequest(r *http.Request, v interface{}) error {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0, post-check=0, pre-check=0")
	r.Header.Set("Strict-Transport-Security", "max-age=31536000 ; includeSubDomains")
	r.SetBasicAuth(fs.Auth.APIKey, "x")
	res, err := fs.client.Do(r)
	if err != nil {
		return fmt.Errorf("error making %s request to %s", r.Method, r.URL)
	}
	defer res.Body.Close()

	// If status code is not ok attempt to read the response in plain text
	if res.StatusCode != 200 && res.StatusCode != 201 {
		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("API request error: %s. unable to retrieve plain text response: %s", res.Status, err.Error())
		}
		return fmt.Errorf("API request error: %s", string(responseData))
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return fmt.Errorf("API request was successful but error occured error decoding response body")
	}

	return nil
}
