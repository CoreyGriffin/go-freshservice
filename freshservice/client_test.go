package freshservice_test

import (
	"context"
	"testing"

	"github.com/CoreyGriffin/go-freshservice/freshservice"
	"github.com/stretchr/testify/assert"
)

var (
	domain = "https://domain.freshservice.com"
	apiKey = "testAPIKey"
)

func TestNewClientDefaultHTTP(t *testing.T) {
	c, err := freshservice.New(nil, domain, apiKey, nil)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, context.Background(), c.Context)
}

func TestNewClientSuccess(t *testing.T) {
	c, err := freshservice.New(nil, domain, apiKey, nil)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, context.Background(), c.Context)
	assert.Equal(t, apiKey, c.Auth.APIKey)
}

func TestNewClientFailMissingDomain(t *testing.T) {
	_, err := freshservice.New(nil, "", apiKey, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "A valid Freshservice domain is required to create a new API client", err.Error())
}

func TestNewClientFailMissingAPIKey(t *testing.T) {
	_, err := freshservice.New(nil, domain, "", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "A valid Freshservice API key is required to create a new API client", err.Error())
}
