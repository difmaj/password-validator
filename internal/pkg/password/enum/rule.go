package enum

// Rule is the enum that defines the rules of the password.
type Rule string

// Constants that defines the rules of the password.
const (
	MinSize         Rule = "minSize"
	MinUppercase    Rule = "minUppercase"
	MinLowercase    Rule = "minLowercase"
	MinDigit        Rule = "minDigit"
	MinSpecialChars Rule = "minSpecialChars"
	NoRepeted       Rule = "noRepeted"
)

// String returns the string representation of the rule.
func (r Rule) String() string {
	return string(r)
}
