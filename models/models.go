package models

// Login ...
type Login struct {
	Email     string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
}
