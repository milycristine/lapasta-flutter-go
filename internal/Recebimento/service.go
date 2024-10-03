package recebimento

import (
    "lapasta/internal/models"
)

type RecebimentoService interface {
    CriarRecebimento(recebimento *models.Recebimento) error
    ListarRecebimentos() ([]models.Recebimento, error)
}

type recebimentoService struct {
    repo RecebimentoRepository
}

func NovoRecebimentoService(repo RecebimentoRepository) RecebimentoService {
    return &recebimentoService{repo: repo}
}

func (s *recebimentoService) CriarRecebimento(recebimento *models.Recebimento) error {
    return s.repo.CriarRecebimento(recebimento)
}

func (s *recebimentoService) ListarRecebimentos() ([]models.Recebimento, error) {
    return s.repo.ListarRecebimentos()
}
