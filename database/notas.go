package sql

import (
	"database/sql"
	"fmt"
	"lapasta/internal/models"
)

func (s *SQLStr) CriarNota(nota *models.Nota) error {
	query := `
        INSERT INTO Notas (Tipo, Valor, IdFuncionario, Url_Imagem, dia) 
        VALUES (@Tipo, @Valor, @IdFuncionario, @UrlImagem, @Dia)
    `
	_, err := s.db.Exec(query,
		sql.Named("Tipo", nota.Tipo),
		sql.Named("Valor", nota.Valor),
		sql.Named("IdFuncionario", nota.IdFuncionario),
		sql.Named("UrlImagem", nota.UrlImagem),
		sql.Named("Dia", nota.Dia),
	)

	if err != nil {
		return fmt.Errorf("erro ao inserir nota: %w", err)
	}
	return nil
}

func (s *SQLStr) ListarNotas() ([]models.Nota, error) {
	var notas []models.Nota

	query := `
        SELECT 
            n.Id, n.Tipo, n.Valor, n.IdFuncionario, n.Url_Imagem, n.dia, 
            f.Nome AS nomeFuncionario 
        FROM Notas n
        JOIN Funcionarios f ON n.IdFuncionario = f.Id
    `
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar notas: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var n models.Nota
		if err := rows.Scan(&n.Id, &n.Tipo, &n.Valor, &n.IdFuncionario, &n.UrlImagem, &n.Dia, &n.NomeFuncionario); err != nil {
			return nil, fmt.Errorf("erro ao escanear nota: %w", err)
		}
		notas = append(notas, n)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das linhas: %w", err)
	}

	return notas, nil
}
