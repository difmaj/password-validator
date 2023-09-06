package password

import (
	"github.com/difmaj/password-validator/internal/pkg/entities"
	"github.com/difmaj/password-validator/internal/pkg/entities/enums"
)

// Verify checks if the password is valid according to the given rules.
func (r *Repository) Verify(request *entities.Request) (*entities.Response, error) {

	noMatch := make([]enums.Rule, 0)
	for _, rule := range request.Rules {
		rules, exist := r.rules[rule.Rule]
		if !exist {
			return nil, ErrRuleNotFound
		}

		verify, err := rules(request.Password, rule.Value)
		if err != nil {
			return nil, err
		}

		if !verify {
			noMatch = append(noMatch, rule.Rule)
		}
	}

	return &entities.Response{
		Verify:  len(noMatch) == 0,
		NoMatch: &noMatch,
	}, nil
}
