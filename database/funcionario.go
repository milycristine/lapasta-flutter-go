package sql

import (
	"database/sql"
	"fmt"
	"lapasta/internal/models"
)


func (r *SQLStr) CriarFuncionario(funcionario *models.Funcionario) error {
	query := `
		INSERT INTO Funcionarios (Nome, Sobrenome, Cpf, Rg, DataNasc, Email, Senha, Cargo, DateAdmissao, HoraEntrada, HoraSaida) 
		VALUES (@Nome, @Sobrenome, @Cpf, @Rg, @DataNasc, @Email, @Senha, @Cargo, @DateAdmissao, @HoraEntrada, @HoraSaida)
	`
	_, err := r.db.Exec(query,
		sql.Named("Nome", funcionario.Nome),
		sql.Named("Sobrenome", funcionario.Sobrenome),
		sql.Named("Cpf", funcionario.Cpf),
		sql.Named("Rg", funcionario.Rg),
		sql.Named("DataNasc", funcionario.DataNasc),
		sql.Named("Email", funcionario.Email),
		sql.Named("Senha", funcionario.Senha),
		sql.Named("Cargo", funcionario.Cargo),
		sql.Named("DateAdmissao", funcionario.DateAdmissao),
		sql.Named("HoraEntrada", funcionario.HoraEntrada),
		sql.Named("HoraSaida", funcionario.HoraSaida),
	)

	if err != nil {
		return fmt.Errorf("erro ao inserir funcionário: %w", err)
	}
	return nil
}

func (r *SQLStr) ListarFuncionarios() ([]models.Funcionario, error) {
	var funcionarios []models.Funcionario

	query := "SELECT Id, Nome, Sobrenome, Cpf, Rg, DataNasc, Email, Cargo, DateAdmissao, HoraEntrada, HoraSaida FROM Funcionarios"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar funcionários: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var f models.Funcionario
		if err := rows.Scan(&f.Id, &f.Nome, &f.Sobrenome, &f.Cpf, &f.Rg, &f.DataNasc, &f.Email, &f.Cargo, &f.DateAdmissao, &f.HoraEntrada, &f.HoraSaida); err != nil {
			return nil, fmt.Errorf("erro ao escanear funcionário: %w", err)
		}
		funcionarios = append(funcionarios, f)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das linhas: %w", err)
	}

	return funcionarios, nil
}
