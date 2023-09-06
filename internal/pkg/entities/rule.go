package entities

import "github.com/difmaj/password-validator/internal/pkg/entities/enums"

// Rule is the entity that defines the rules of the password.
type Rule struct {
	// Rule is the name of the rule.
	Rule enums.Rule `json:"rule"`
	// Value is the value of the rule.
	Value int `json:"value"`
}
