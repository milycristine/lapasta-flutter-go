package nota

import (
	database "lapasta/database"
	"lapasta/internal/models"
)

type NotaRepository interface {
	CriarNota(nota *models.Nota) error
	ListarNotas() ([]models.Nota, error)
}

type notaRepository struct {
	db *database.SQLStr
}

func NovoNotaRepository(db *database.SQLStr) NotaRepository {
	return &notaRepository{
		db: db,
	}
}

func (r *notaRepository) CriarNota(nota *models.Nota) error {
	return r.db.CriarNota(nota)
}

func (r *notaRepository) ListarNotas() ([]models.Nota, error) {
	return r.db.ListarNotas()
}
