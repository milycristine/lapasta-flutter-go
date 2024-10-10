package documento

import (
	database "lapasta/database"
	"lapasta/internal/models"
)

type DocumentoRepository interface {
	CriarDocumento(documento *models.Documento) error
	ListarDocumentos() ([]models.Documento, error)
}

type documentoRepository struct {
	db *database.SQLStr
}

func NovoDocumentoRepository(db *database.SQLStr) DocumentoRepository {
	return &documentoRepository{
		db: db,
	}
}

func (r *documentoRepository) CriarDocumento(documento *models.Documento) error {
	return r.db.CriarDocumento(documento)
}

func (r *documentoRepository) ListarDocumentos() ([]models.Documento, error) {
	return r.db.ListarDocumentos()
}
