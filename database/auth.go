package sql

import (
	"database/sql"
	"fmt"
	"lapasta/internal/models"
)


func (s *SQLStr) Autenticar(username string) (models.Login, error) {
	var Auth models.Login

	query := "SELECT * FROM AUTH WHERE username = '%s'"
	err := s.db.QueryRow(fmt.Sprintf(query, username)).Scan(&Auth.Username, &Auth.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return Auth, fmt.Errorf("usuário não encontrado")
		}
		return Auth, err
	}

	return Auth, nil
}