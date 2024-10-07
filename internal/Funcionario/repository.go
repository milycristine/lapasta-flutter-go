package funcionario

import (
	database "lapasta/database" // Ajuste conforme necessário
	"lapasta/internal/models"
)

// FuncionarioRepository interface para definir métodos do repositório.
type FuncionarioRepository interface {
	CriarFuncionario(funcionario *models.Funcionario) error
	ListarFuncionarios() ([]models.Funcionario, error)
}

// funcionarioRepository implementação da interface FuncionarioRepository.
type funcionarioRepository struct {
	db *database.SQLStr
}

// NovoFuncionarioRepository cria uma nova instância de funcionarioRepository.
func NovoFuncionarioRepository(db *database.SQLStr) FuncionarioRepository {
	return &funcionarioRepository{
		db: db,
	}
}

// CriarFuncionario utiliza a implementação existente da database.
func (r *funcionarioRepository) CriarFuncionario(funcionario *models.Funcionario) error {
	return r.db.CriarFuncionario(funcionario)
}

// ListarFuncionarios utiliza a implementação existente da database.
func (r *funcionarioRepository) ListarFuncionarios() ([]models.Funcionario, error) {
	return r.db.ListarFuncionarios()
}
