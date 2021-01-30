package freshservice

import (
	"errors"
	"fmt"
)

// ErrorResponse represents a Freshservice API error
type ErrorResponse struct {
	Description string  `json:"description"`
	Errors      []Error `json:"errors"`
}

// Error holds the details of a Freshservice error
type Error struct {
	Field   string `json:"field"` // Applicable to HTTP 400 errors only.
	Message string `json:"message"`
	Code    string `json:"code"`
}

// Helper to be used for API client config errors
func missingClientConfigErr(attr string) error {
	errTxt := fmt.Sprintf("A valid Freshservice %s is required to create a new API client", attr)
	return errors.New(errTxt)
}
