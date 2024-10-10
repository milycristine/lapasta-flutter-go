package nota

import (
	"lapasta/internal/models"
)

type NotaService interface {
	CriarNota(nota *models.Nota) error
	ListarNotas() ([]models.Nota, error)
}

type notaService struct {
	repo NotaRepository
}

func NovaNotaService(repo NotaRepository) NotaService {
	return &notaService{
		repo: repo,
	}
}

func (s *notaService) CriarNota(nota *models.Nota) error {
	return s.repo.CriarNota(nota)
}

func (s *notaService) ListarNotas() ([]models.Nota, error) {
	return s.repo.ListarNotas()
}
