package funcionario

import (
	"lapasta/internal/models"
)

type FuncionarioService interface {
	CriarFuncionario(funcionario *models.Funcionario) error
	ListarFuncionarios() ([]models.Funcionario, error)
}

type funcionarioService struct {
	repo FuncionarioRepository
}

func NovoFuncionarioService(repo FuncionarioRepository) FuncionarioService {
	return &funcionarioService{
		repo: repo,
	}
}


func (s *funcionarioService) CriarFuncionario(funcionario *models.Funcionario) error {
	return s.repo.CriarFuncionario(funcionario)
}

func (s *funcionarioService) ListarFuncionarios() ([]models.Funcionario, error) {
	return s.repo.ListarFuncionarios()
}
