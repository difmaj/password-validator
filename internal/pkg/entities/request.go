package entities

// Rule is the entity that defines the rules of the password.
type Request struct {
	// Password is the password to be verified.
	Password string `json:"password"`
	// Rules is the rules that the password must follow.
	Rules []*Rule `json:"rules"`
}
