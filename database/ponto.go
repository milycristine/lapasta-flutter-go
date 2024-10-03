package sql

import (
	"database/sql"
	"fmt"
	"lapasta/internal/models"
)



func (s *SQLStr) CriarPonto(ponto *models.Ponto) error {
	query := `
        INSERT INTO Ponto (HManha, HAlmocoEntrada, HAlmocoSaida, HNoite, Dia, Situacao, IdFuncionario) 
        VALUES (@HManha, @HAlmocoEntrada, @HAlmocoSaida, @HNoite, @Dia, @Situacao, @IdFuncionario)
    `
	_, err := s.db.Exec(query,
		sql.Named("HManha", ponto.HManha),
		sql.Named("HAlmocoEntrada", ponto.HAlmocoEntrada),
		sql.Named("HAlmocoSaida", ponto.HAlmocoSaida),
		sql.Named("HNoite", ponto.HNoite),
		sql.Named("Dia", ponto.Dia),
		sql.Named("Situacao", ponto.Situacao),
		sql.Named("IdFuncionario", ponto.IdFuncionario))

	if err != nil {
		return fmt.Errorf("erro ao inserir ponto: %w", err)
	}
	return nil
}
func (s *SQLStr) ListarPontos() ([]models.Ponto, error) {
	var pontos []models.Ponto

	query := "SELECT Id, HManha, HAlmocoEntrada, HAlmocoSaida, HNoite, Dia, Situacao, IdFuncionario FROM Ponto"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar pontos: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Ponto
		if err := rows.Scan(&p.Id, &p.HManha, &p.HAlmocoEntrada, &p.HAlmocoSaida, &p.HNoite, &p.Dia, &p.Situacao, &p.IdFuncionario); err != nil {
			return nil, fmt.Errorf("erro ao escanear ponto: %w", err)
		}
		pontos = append(pontos, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das linhas: %w", err)
	}

	return pontos, nil
}
