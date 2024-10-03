// repository/ponto.go
package ponto

import (
	database "lapasta/database"
	"lapasta/internal/models"
)

type PontoRepository interface {
	CriarPonto(ponto *models.Ponto) error
	ListarPontos() ([]models.Ponto, error)
}

type pontoRepository struct {
	db *database.SQLStr
}

func NovoPontoRepository(db *database.SQLStr) PontoRepository {
	return &pontoRepository{
		db: db,
	}
}

func (r *pontoRepository) CriarPonto(ponto *models.Ponto) error {
	return r.db.CriarPonto(ponto)
}

func (r *pontoRepository) ListarPontos() ([]models.Ponto, error) {
	return r.db.ListarPontos()
}
