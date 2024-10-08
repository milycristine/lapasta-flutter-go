package funcionario

import (
	database "lapasta/database" 
	"lapasta/internal/models"
)

type FuncionarioRepository interface {
	CriarFuncionario(funcionario *models.Funcionario) error
	ListarFuncionarios() ([]models.Funcionario, error)
}

type funcionarioRepository struct {
	db *database.SQLStr
}

func NovoFuncionarioRepository(db *database.SQLStr) FuncionarioRepository {
	return &funcionarioRepository{
		db: db,
	}
}

func (r *funcionarioRepository) CriarFuncionario(funcionario *models.Funcionario) error {
	return r.db.CriarFuncionario(funcionario)
}

func (r *funcionarioRepository) ListarFuncionarios() ([]models.Funcionario, error) {
	return r.db.ListarFuncionarios()
}
