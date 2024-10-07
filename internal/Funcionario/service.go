package funcionario

import (
	"lapasta/internal/models"
)

// FuncionarioService interface para definir métodos do serviço.
type FuncionarioService interface {
	CriarFuncionario(funcionario *models.Funcionario) error
	ListarFuncionarios() ([]models.Funcionario, error)
}

// funcionarioService implementação da interface FuncionarioService.
type funcionarioService struct {
	repo FuncionarioRepository
}

// NovoFuncionarioService cria uma nova instância de funcionarioService.
func NovoFuncionarioService(repo FuncionarioRepository) FuncionarioService {
	return &funcionarioService{
		repo: repo,
	}
}

// CriarFuncionario
func (s *funcionarioService) CriarFuncionario(funcionario *models.Funcionario) error {
	return s.repo.CriarFuncionario(funcionario)
}

// ListarFuncionarios
func (s *funcionarioService) ListarFuncionarios() ([]models.Funcionario, error) {
	return s.repo.ListarFuncionarios()
}
