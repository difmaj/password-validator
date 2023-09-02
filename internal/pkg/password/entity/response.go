package entity

import "github.com/difmaj/password-validator-challenge/internal/pkg/password/enum"

// Response is the entity that defines the response of the service.
type Response struct {
	// Verify is the result of the password verification.
	Verify bool `json:"verify"`
	// NoMatch is the list of rules that the password does not match.
	NoMatch *[]enum.Rule `json:"noMatch,omitempty"`
}
