package sql

import (
	"database/sql"
	"fmt"
	"lapasta/internal/models"
)

func (s *SQLStr) CriarRecebimento(recebimento *models.Recebimento) error {
	query := `
        INSERT INTO Recebimento (Dia, Empresa, Produto, UrlImagem, IdResponsavel) 
        VALUES (@Dia, @Empresa, @Produto, @UrlImagem, @IdResponsavel)
    `
	_, err := s.db.Exec(query,
		sql.Named("Dia", recebimento.Dia),
		sql.Named("Empresa", recebimento.Empresa),
		sql.Named("Produto", recebimento.Produto),
		sql.Named("UrlImagem", recebimento.UrlImagem),
		sql.Named("IdResponsavel", recebimento.IdResponsavel))

	if err != nil {
		return fmt.Errorf("erro ao inserir recebimento: %w", err)
	}
	return nil
}

func (s *SQLStr) ListarRecebimentos() ([]models.Recebimento, error) {
	var recebimentos []models.Recebimento

	query := "SELECT Id, Dia, Empresa, Produto, UrlImagem, IdResponsavel FROM Recebimento"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar recebimentos: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Recebimento
		if err := rows.Scan(&r.Id, &r.Dia, &r.Empresa, &r.Produto, &r.UrlImagem, &r.IdResponsavel); err != nil {
			return nil, fmt.Errorf("erro ao escanear recebimento: %w", err)
		}
		recebimentos = append(recebimentos, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das linhas: %w", err)
	}

	return recebimentos, nil
}
