package password

import (
	"fmt"

	"github.com/difmaj/password-validator-challenge/internal/pkg/password/entity"
)

// Service is the struct that implements the UseCase interface.
type Service struct {
	repository Repository
}

// New returns a new Service instance.
func New(repository Repository) UseCase {
	return &Service{repository: repository}
}

func (s *Service) Verify(request *entity.Request) (*entity.Response, error) {
	response, err := s.repository.Verify(request)
	if err != nil {
		return nil, fmt.Errorf("password.Verify: %w", err)
	}
	return response, nil
}
