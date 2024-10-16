package pagamento

import (
	database "lapasta/database"
	"lapasta/internal/models"
)

type PagamentoRepository interface {
	CriarPagamento(pagamento *models.Pagamento) error
	ListarPagamentos() ([]models.Pagamento, error)
	ListarPagamentosPorDia(dia int) ([]models.Pagamento, error)
	AtualizarStatusPagamento(id int, statusId int) error
}

type pagamentoRepository struct {
	db *database.SQLStr
}

func NovoPagamentoRepository(db *database.SQLStr) PagamentoRepository {
	return &pagamentoRepository{
		db: db,
	}
}

func (r *pagamentoRepository) CriarPagamento(pagamento *models.Pagamento) error {
	return r.db.CriarPagamento(pagamento)
}

func (r *pagamentoRepository) ListarPagamentos() ([]models.Pagamento, error) {
	return r.db.ListarPagamentos()
}

func (r *pagamentoRepository) ListarPagamentosPorDia(dia int) ([]models.Pagamento, error) {
	return r.db.ListarPagamentosPorDia(dia)
}
func (r *pagamentoRepository) AtualizarStatusPagamento(id int, statusId int) error {
	return r.db.AtualizarStatusPagamento(id, statusId)
}
