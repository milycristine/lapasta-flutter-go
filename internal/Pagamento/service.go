package pagamento

import (
	"lapasta/internal/models"
)

type PagamentoService interface {
	CriarPagamento(pagamento *models.Pagamento) error
	ListarPagamentos() ([]models.Pagamento, error)
	ListarPagamentosPorDia(dia int) ([]models.Pagamento, error)
	AtualizarStatusPagamento(id int, statusId int) error
}

type pagamentoService struct {
	repo PagamentoRepository
}

func NovoPagamentoService(repo PagamentoRepository) PagamentoService {
	return &pagamentoService{
		repo: repo,
	}
}

func (s *pagamentoService) CriarPagamento(pagamento *models.Pagamento) error {
	return s.repo.CriarPagamento(pagamento)
}

func (s *pagamentoService) ListarPagamentos() ([]models.Pagamento, error) {
	return s.repo.ListarPagamentos()
}
func (s *pagamentoService) ListarPagamentosPorDia(dia int) ([]models.Pagamento, error) {
	return s.repo.ListarPagamentosPorDia(dia)
}
func (s *pagamentoService) AtualizarStatusPagamento(id int, statusId int) error {
	return s.repo.AtualizarStatusPagamento(id, statusId)
}
