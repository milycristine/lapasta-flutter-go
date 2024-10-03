// internal/auth/repository.go
package auth

import (
	database "lapasta/database"
	"lapasta/internal/models"
)

// AuthRepository interface para o repositório de autenticação.
type AuthRepository interface {
	Autenticar(username string) (models.Login, error)
}

// authRepository é a implementação do AuthRepository.
type authRepository struct {
	db *database.SQLStr // Conexão com o banco de dados
}

// NewAuthRepository cria uma nova instância de authRepository.
func NewAuthRepository(db *database.SQLStr) AuthRepository {
	return &authRepository{db: db}
}

// Autenticar busca as informações do usuário no banco de dados.
func (r *authRepository) Autenticar(username string) (models.Login, error) {
	// Utiliza o método Autenticar do repositório SQL
	login, err := r.db.Autenticar(username)
	if err != nil {
		return models.Login{}, err
	}
	return login, nil
}
