package password

import (
	"unicode"

	"github.com/difmaj/password-validator/internal/pkg/entities/enums"
)

// Rule is the type that defines the rules of the password.
type Rule func(string, int) (bool, error)

// rules is the map of rules.
var rules = map[enums.Rule]Rule{
	enums.MinSize:         MinSize,
	enums.MinUppercase:    MinUppercase,
	enums.MinLowercase:    MinLowercase,
	enums.MinDigit:        MinDigit,
	enums.MinSpecialChars: MinSpecialChars,
	enums.NoRepeted:       NoRepeted,
}

// MinSize check that pwd has at least {Rule.Value) characters.
func MinSize(pwd string, value int) (bool, error) {
	return len(pwd) >= value, nil
}

// MinUppercase check that pwd has at least {Rule.Value) uppercase characters.
func MinUppercase(pwd string, value int) (bool, error) {
	var count int
	for _, r := range pwd {
		if unicode.IsUpper(r) {
			count++
		}
	}
	return count >= value, nil
}

// MinLowercase check that pwd has at least {Rule.Value) lowercase characters.
func MinLowercase(pwd string, value int) (bool, error) {
	var count int
	for _, r := range pwd {
		if unicode.IsLower(r) {
			count++
		}
	}
	return count >= value, nil
}

// MinDigit check that pwd has at least {Rule.Value) digits.
func MinDigit(pwd string, value int) (bool, error) {
	var count int
	for _, r := range pwd {
		if unicode.IsDigit(r) {
			count++
		}
	}
	return count >= value, nil
}

// MinSpecialChars check that pwd has at least {Rule.Value) special characters.
func MinSpecialChars(pwd string, value int) (bool, error) {
	var count int
	for _, r := range pwd {
		if unicode.IsPunct(r) {
			count++
		}
	}
	return count >= value, nil
}

// NoRepeted check that pwd has no repeated characters.
func NoRepeted(pwd string, value int) (bool, error) {
	var count int
	for i, r := range pwd {
		for j, r2 := range pwd {
			if i != j && r == r2 {
				count++
			}
		}
	}
	return count == 0, nil
}
