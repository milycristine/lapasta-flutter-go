package auth

import (
	database "lapasta/database"
	"lapasta/internal/models"
)

type AuthRepository interface {
	Autenticar(username string) (models.Login, error)
}

type authRepository struct {
	db *database.SQLStr
}

func NewAuthRepository(db *database.SQLStr) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Autenticar(username string) (models.Login, error) {
	login, err := r.db.Autenticar(username)
	if err != nil {
		return models.Login{}, err
	}
	return login, nil
}
