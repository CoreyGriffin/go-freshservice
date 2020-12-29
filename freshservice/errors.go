package freshservice

import (
	"errors"
	"fmt"
)

// Helper to be used for API client config errors
func missingClientConfigErr(attr string) error {
	errTxt := fmt.Sprintf("a valid Freshservice %s is required to create a new API client", attr)
	return errors.New(errTxt)
}
