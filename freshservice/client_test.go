package freshservice_test

import (
	"context"
	"testing"

	"github.com/CoreyGriffin/go-freshservice/freshservice"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	domain   = "https://domain.freshservice.com"
	username = "testUser"
	apiKey   = "testAPIKey"
	logger   = logrus.New()
)

func TestNewClientDefaultHTTP(t *testing.T) {
	c, err := freshservice.New(nil, domain, 1, username, apiKey, logger, nil)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, context.Background(), c.Context)
}

func TestNewClientSuccess_v1(t *testing.T) {
	c, err := freshservice.New(nil, domain, 1, username, apiKey, logger, nil)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, context.Background(), c.Context)
	assert.Equal(t, 1, c.Version)
	assert.Equal(t, username, c.Auth.Username)
	assert.Equal(t, apiKey, c.Auth.APIKey)
	assert.Equal(t, "https://domain.freshservice.com", c.Endpoint)
}

func TestNewClientSuccess_v2(t *testing.T) {
	c, err := freshservice.New(nil, domain, 2, username, apiKey, logger, nil)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, context.Background(), c.Context)
	assert.Equal(t, 2, c.Version)
	assert.Equal(t, username, c.Auth.Username)
	assert.Equal(t, apiKey, c.Auth.APIKey)
	assert.Equal(t, "https://domain.freshservice.com/api/v2", c.Endpoint)
}

func TestNewClientFailMissingDomain(t *testing.T) {
	_, err := freshservice.New(nil, "", 2, username, apiKey, logger, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "a valid Freshservice domain is required to create a new API client", err.Error())
}

func TestNewClientFailMissingUsername(t *testing.T) {
	_, err := freshservice.New(nil, domain, 2, "", apiKey, logger, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "a valid Freshservice username is required to create a new API client", err.Error())
}

func TestNewClientFailMissingAPIKey(t *testing.T) {
	_, err := freshservice.New(nil, domain, 2, username, "", logger, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "a valid Freshservice API Key is required to create a new API client", err.Error())
}
