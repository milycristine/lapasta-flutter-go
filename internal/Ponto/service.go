// service/ponto.go
package ponto

import (
	"lapasta/internal/models"
)

type PontoService interface {
	CriarPonto(ponto *models.Ponto) error
	ListarPontos() ([]models.Ponto, error)
} 

type pontoService struct {
	repo PontoRepository
}

func NovoPontoService(repo PontoRepository) PontoService {
	return &pontoService{
		repo: repo,
	}
}

func (s *pontoService) CriarPonto(ponto *models.Ponto) error {
	// Validações e regras de negócio podem ser aplicadas aqui
	return s.repo.CriarPonto(ponto)
}

func (s *pontoService) ListarPontos() ([]models.Ponto, error) {
	return s.repo.ListarPontos()
}
