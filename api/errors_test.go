package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingClientConfigErr(t *testing.T) {
	cases := []struct {
		Attribute string
		Expected  string
	}{
		{
			Attribute: "testing",
			Expected:  "a valid Freshservice testing is required to create a new API client",
		},
		{
			Attribute: "API Key",
			Expected:  "a valid Freshservice API Key is required to create a new API client",
		},
		{
			Attribute: "username",
			Expected:  "a valid Freshservice username is required to create a new API client",
		},
		{
			Attribute: "domain",
			Expected:  "a valid Freshservice domain is required to create a new API client",
		},
	}

	for _, c := range cases {
		err := missingClientConfigErr(c.Attribute)
		assert.NotNil(t, err)
		assert.Equal(t, c.Expected, err.Error())
	}
}
