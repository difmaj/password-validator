package repository

import (
	"errors"

	"github.com/difmaj/password-validator-challenge/internal/pkg/password/entity"
)

func (r *Repository) Verify(request *entity.Request) (*entity.Response, error) {
	return nil, errors.New("not implemented")
}
