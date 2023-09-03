package password

import "github.com/difmaj/password-validator/internal/pkg/password/entity"

// UseCase is the interface that defines a service methods.
type UseCase interface {
	Verify(*entity.Request) (*entity.Response, error)
}
