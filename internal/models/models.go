package models

type Login struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// Response ...
type Response struct {
	Status string
	Error  string
	Data   interface{}
}
