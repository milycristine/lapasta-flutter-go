package auth

import (
	"fmt"
	"lapasta/internal/models"
)

type AuthService interface {
	Login(username, password string) (models.Login, error)
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(username, password string) (models.Login, error) {
	login, err := s.repo.Autenticar(username)
	if err != nil {
		return models.Login{}, err
	}

	if login.Password != password {
		return models.Login{}, fmt.Errorf("nome ou senha inválidos")
	}
	return login, nil
}
