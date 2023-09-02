package repository

import (
	"github.com/difmaj/password-validator-challenge/internal/pkg/password"
	"github.com/difmaj/password-validator-challenge/internal/pkg/password/enum"
)

// Repository is the struct that implements the password.Repository interface.
type Repository struct {
	rules map[enum.Rule]Rule
}

// New returns a new Repository instance.
func New() password.Repository {
	return &Repository{
		rules: rules,
	}
}
