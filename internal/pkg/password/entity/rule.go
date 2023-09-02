package entity

import "github.com/difmaj/password-validator-challenge/internal/pkg/password/enum"

// Rule is the entity that defines the rules of the password.
type Rule struct {
	// Rule is the name of the rule.
	Rule enum.Rule `json:"rule"`
	// Value is the value of the rule.
	Value int `json:"value"`
}
