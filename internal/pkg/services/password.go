package services

import (
	"fmt"

	"github.com/difmaj/password-validator/internal/pkg/entities"
)

// PasswordRepository is the interface that wraps the Verify method.
type PasswordRepository interface {
	// Verify checks if the password is valid according to the given rules.
	Verify(*entities.Request) (*entities.Response, error)
}

// PasswordService is the struct that implements the UseCase interface.
type PasswordService struct {
	repository PasswordRepository
}

// NewPasswordService returns a new Service instance.
func NewPasswordService(repository PasswordRepository) *PasswordService {
	return &PasswordService{repository: repository}
}

// Verify checks if the password is valid according to the given rules.
func (s *PasswordService) Verify(request *entities.Request) (*entities.Response, error) {
	response, err := s.repository.Verify(request)
	if err != nil {
		return nil, fmt.Errorf("password.Verify: %w", err)
	}
	return response, nil
}
