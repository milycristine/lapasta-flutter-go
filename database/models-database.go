package sql

import (
	"database/sql"
	"net/url"
)

// SQLStr ...
type SQLStr struct {
	url *url.URL
	db  *sql.DB
}

// Login representa os dados necessários para autenticar um usuário.
type Login struct {
	Email string `json:"email,omitempty"` // E-mail do usuário
	Senha string `json:"senha,omitempty"` // Senha do usuário
}
