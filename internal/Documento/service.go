package documento

import (
	"lapasta/internal/models"
)

type DocumentoService interface {
	CriarDocumento(documento *models.Documento) error
	ListarDocumentos() ([]models.Documento, error)
}

type documentoService struct {
	repo DocumentoRepository
}

func NovoDocumentoService(repo DocumentoRepository) DocumentoService {
	return &documentoService{
		repo: repo,
	}
}

func (s *documentoService) CriarDocumento(documento *models.Documento) error {
	return s.repo.CriarDocumento(documento)
}

func (s *documentoService) ListarDocumentos() ([]models.Documento, error) {
	return s.repo.ListarDocumentos()
}
