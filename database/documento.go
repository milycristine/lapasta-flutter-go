package sql

import (
	"database/sql"
	"fmt"
	"lapasta/internal/models"
)

func (s *SQLStr) CriarDocumento(doc *models.Documento) error {
	query := `
        INSERT INTO Documentos (Titulo, Url, Data_Criacao, IdFuncionario) 
        VALUES (@Titulo, @Url, @Data_Criacao, @IdFuncionario)
    `
	_, err := s.db.Exec(query,
		sql.Named("Titulo", doc.Titulo),
		sql.Named("Url", doc.Url),
		sql.Named("Data_Criacao", doc.DataCriacao),
		sql.Named("IdFuncionario", doc.IdFuncionario),
	)

	if err != nil {
		return fmt.Errorf("erro ao inserir documento: %w", err)
	}
	return nil
}

func (s *SQLStr) ListarDocumentos() ([]models.Documento, error) {
	var docs []models.Documento

	query := "SELECT Id, Titulo, Url, Data_Criacao, IdFuncionario FROM Documentos"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar documentos: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var d models.Documento
		if err := rows.Scan(&d.Id, &d.Titulo, &d.Url, &d.DataCriacao, &d.IdFuncionario); err != nil {
			return nil, fmt.Errorf("erro ao escanear documento: %w", err)
		}
		docs = append(docs, d)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das linhas: %w", err)
	}

	return docs, nil
}
