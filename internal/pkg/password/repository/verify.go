package repository

import (
	"github.com/difmaj/password-validator/internal/pkg/password/entity"
	"github.com/difmaj/password-validator/internal/pkg/password/enum"
)

func (r *Repository) Verify(request *entity.Request) (*entity.Response, error) {

	noMatch := make([]enum.Rule, 0)
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

	return &entity.Response{
		Verify:  len(noMatch) == 0,
		NoMatch: &noMatch,
	}, nil
}
