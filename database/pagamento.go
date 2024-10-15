package sql

import (
	"database/sql"
	"fmt"
	"lapasta/internal/models"
)

func (s *SQLStr) CriarPagamento(pagamento *models.Pagamento) error {
	query := `
        INSERT INTO Pagamentos (FuncionarioId, DataPagamento, Valor, StatusId) 
        VALUES (@FuncionarioId, NULL, @Valor, 2)  -- Inicializa como Pendente
    `
	_, err := s.db.Exec(query,
		sql.Named("FuncionarioId", pagamento.FuncionarioId),
		sql.Named("Valor", pagamento.Valor))

	if err != nil {
		return fmt.Errorf("erro ao inserir pagamento: %w", err)
	}
	return nil
}

func (s *SQLStr) ListarPagamentos() ([]models.Pagamento, error) {
	var pagamentos []models.Pagamento

	query := `
        SELECT 
            p.Id, 
            p.FuncionarioId, 
            p.DataPagamento, 
            p.Valor, 
            p.StatusId, 
            f.Cpf, 
            f.Nome 
        FROM 
            Pagamentos p
        JOIN 
            Funcionarios f ON p.FuncionarioId = f.Id
    `
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar pagamentos: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Pagamento
		if err := rows.Scan(&p.Id, &p.FuncionarioId, &p.DataPagamento, &p.Valor, &p.StatusId, &p.CpfFuncionario, &p.NomeFuncionario); err != nil {
			return nil, fmt.Errorf("erro ao escanear pagamento: %w", err)
		}
		pagamentos = append(pagamentos, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das linhas: %w", err)
	}

	return pagamentos, nil
}

func (s *SQLStr) ListarPagamentosPorDia(dia int) ([]models.Pagamento, error) {
	var pagamentos []models.Pagamento

	query := `
		SELECT 
            p.Id, 
            p.FuncionarioId, 
            p.DataPagamento, 
            p.Valor, 
            p.StatusId,
            f.Cpf,
            f.Nome 
        FROM 
            Pagamentos p
        JOIN 
            Funcionarios f ON p.FuncionarioId = f.Id
        WHERE 
            DAY(p.DataPagamento) = @DiaPagamento
	`
	rows, err := s.db.Query(query, sql.Named("DiaPagamento", dia))
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar pagamentos para o dia %d: %w", dia, err)
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Pagamento
		if err := rows.Scan(&p.Id, &p.FuncionarioId, &p.DataPagamento, &p.Valor, &p.StatusId, &p.CpfFuncionario, &p.NomeFuncionario); err != nil {
			return nil, fmt.Errorf("erro ao escanear pagamento: %w", err)
		}
		pagamentos = append(pagamentos, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das linhas: %w", err)
	}

	return pagamentos, nil
}

func (s *SQLStr) AtualizarStatusPagamento(id int, statusId int) error {
	query := `
        UPDATE Pagamentos 
        SET StatusId = @StatusId 
        WHERE Id = @Id
    `
	_, err := s.db.Exec(query,
		sql.Named("Id", id),
		sql.Named("StatusId", statusId))

	if err != nil {
		return fmt.Errorf("erro ao atualizar pagamento: %w", err)
	}
	return nil
}
