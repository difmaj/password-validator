package password

import (
	"testing"

	"github.com/difmaj/password-validator/internal/pkg/entities/enums"
)

type testRule map[string]struct {
	password string
	value    int
	want     bool
}

type testRules map[enums.Rule]testRule

func TestRules(t *testing.T) {
	tests := testRules{
		enums.MinSize: {
			"invalid": {
				password: "1234567",
				value:    8,
				want:     false,
			},
			"valid": {
				password: "1234567",
				value:    7,
				want:     true,
			},
		},
		enums.MinUppercase: {
			"invalid": {
				password: "abc",
				value:    1,
				want:     false,
			},
			"valid": {
				password: "Abc",
				value:    1,
				want:     true,
			},
		},
		enums.MinLowercase: {
			"invalid": {
				password: "ABC",
				value:    1,
				want:     false,
			},
			"valid": {
				password: "Abc",
				value:    1,
				want:     true,
			},
		},
		enums.MinDigit: {
			"invalid": {
				password: "abc",
				value:    1,
				want:     false,
			},
			"valid": {
				password: "abc1",
				value:    1,
				want:     true,
			},
		},
		enums.MinSpecialChars: {
			"invalid": {
				password: "abc",
				value:    1,
				want:     false,
			},
			"valid": {
				password: "abc!",
				value:    1,
				want:     true,
			},
		},
		enums.NoRepeted: {
			"invalid": {
				password: "abcabc",
				value:    1,
				want:     false,
			},
			"valid": {
				password: "abc",
				value:    1,
				want:     true,
			},
		},
	}

	for ruleName, testsRule := range tests {
		rule := rules[ruleName]

		t.Run(ruleName.String(), func(t *testing.T) {
			for testName, tt := range testsRule {
				t.Run(testName, func(t *testing.T) {
					got, err := rule(tt.password, tt.value)
					if err != nil {
						t.Fatalf("%s() unexpected error: %v", ruleName, err)
					}
					if got != tt.want {
						t.Errorf("%s() = %v, want %v", ruleName, got, tt.want)
					}
				})
			}
		})
	}
}
