package sql

import (
	"database/sql"
	"fmt"
	"lapasta/internal/models"
	"net/url"

	"golang.org/x/crypto/bcrypt"
)

type SQLStr struct {
	url *url.URL
	db  *sql.DB
}

func (s *SQLStr) Autenticar(username string) (models.Login, error) {
	var Auth models.Login

	query := "SELECT * FROM AUTH WHERE username = '%s'"
	err := s.db.QueryRow(fmt.Sprintf(query, username)).Scan(&Auth.Username, &Auth.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return Auth, fmt.Errorf("usuário não encontrado")
		}
		return Auth, err
	}

	return Auth, nil
}

// CriarFuncionario insere um novo funcionário no banco de dados com a senha hasheada.
func (s *SQLStr) CriarFuncionario(funcionario models.Funcionario) error {
	// Gera o hash da senha antes de armazenar
	hashedSenha, err := bcrypt.GenerateFromPassword([]byte(funcionario.Senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Insere o funcionário no banco
	query := "INSERT INTO Funcionarios (Nome, Sobrenome, Cpf, Rg, DataNasc, Email, Senha, Cargo, DateAdmissao, HoraEntrada, HoraSaida) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = s.db.Exec(query, funcionario.Nome, funcionario.Sobrenome, funcionario.Cpf, funcionario.Rg, funcionario.DataNasc, funcionario.Email, hashedSenha, funcionario.Cargo, funcionario.DateAdmissao, funcionario.HoraEntrada, funcionario.HoraSaida)
	if err != nil {
		return err
	}

	return nil
}

// CriarRecebimento insere um novo recebimento no banco de dados.
func (s *SQLStr) CriarRecebimento(recebimento *models.Recebimento) error {
	query := "INSERT INTO Recebimento (Dia, Empresa, Produto, UrlImagem, IdResponsavel) VALUES (?, ?, ?, ?, ?)"
	_, err := s.db.Exec(query, recebimento.Dia, recebimento.Empresa, recebimento.Produto, recebimento.UrlImagem, recebimento.IdResponsavel)
	if err != nil {
		return fmt.Errorf("erro ao inserir recebimento: %w", err)
	}
	return nil
}

// ListarRecebimentos recupera todos os recebimentos do banco de dados.
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
