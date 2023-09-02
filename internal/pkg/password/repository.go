package password

import "github.com/difmaj/password-validator-challenge/internal/pkg/password/entity"

type Repository interface {
	Verify(*entity.Request) (*entity.Response, error)
}
