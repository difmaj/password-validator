package repository

import "github.com/difmaj/password-validator-challenge/internal/pkg/password"

// Repository is the struct that implements the password.Repository interface.
type Repository struct{}

// New returns a new Repository instance.
func New() password.Repository {
	return &Repository{}
}
