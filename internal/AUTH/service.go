// internal/auth/service.go
package auth

import (
	"fmt"
	"lapasta/internal/models"
)

// AuthService interface para os serviços de autenticação.
type AuthService interface {
	Login(username, password string) (models.Login, error)
}

// authService é a implementação do AuthService.
type authService struct {
	repo AuthRepository
}

// NewAuthService cria uma nova instância de authService.
func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo: repo}
}

// Login autentica um usuário com nome de usuário e senha.
func (s *authService) Login(username, password string) (models.Login, error) {
	// Busca o login do repositório
	login, err := s.repo.Autenticar(username)
	if err != nil {
		return models.Login{}, err
	}

	// Verifica a senha fornecida com o hash armazenado
	if login.Password != password {
		return models.Login{}, fmt.Errorf("nome ou senha inválidos")
	}
	return login, nil
}
