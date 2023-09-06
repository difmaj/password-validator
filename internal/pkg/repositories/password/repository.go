package password

import (
	"github.com/difmaj/password-validator/internal/pkg/entities/enums"
	"github.com/difmaj/password-validator/internal/pkg/services"
)

// Repository is the struct that implements the password.Repository interface.
type Repository struct {
	rules map[enums.Rule]Rule
}

// NewRepository returns a new Repository instance.
func NewRepository() services.PasswordRepository {
	return &Repository{
		rules: rules,
	}
}
