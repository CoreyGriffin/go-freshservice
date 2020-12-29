package freshservice

import (
	"context"
	"fmt"
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

// New returns a new Freshservice API client
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
