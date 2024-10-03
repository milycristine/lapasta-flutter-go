package funcionario

import (
	"encoding/json"
	"lapasta/internal/models"
	"net/http"

)

// FuncionarioHandler interface para definir métodos do handler.
type FuncionarioHandler interface {
	CriarFuncionario(w http.ResponseWriter, r *http.Request)
	ListarFuncionarios(w http.ResponseWriter, r *http.Request)
}

// funcionarioHandler implementação da interface FuncionarioHandler.
type funcionarioHandler struct {
	service FuncionarioService
}

// NovoFuncionarioHandler cria uma nova instância de funcionarioHandler.
func NovoFuncionarioHandler(service FuncionarioService) FuncionarioHandler {
	return &funcionarioHandler{
		service: service,
	}
}

// CriarFuncionario
func (h *funcionarioHandler) CriarFuncionario(w http.ResponseWriter, r *http.Request) {
	var funcionario models.Funcionario
	if err := json.NewDecoder(r.Body).Decode(&funcionario); err != nil {
		http.Error(w, "Erro ao decodificar o funcionário", http.StatusBadRequest)
		return
	}

	if err := h.service.CriarFuncionario(&funcionario); err != nil {
		http.Error(w, "Erro ao criar funcionário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// ListarFuncionarios
func (h *funcionarioHandler) ListarFuncionarios(w http.ResponseWriter, r *http.Request) {
	funcionarios, err := h.service.ListarFuncionarios()
	if err != nil {
		http.Error(w, "Erro ao listar funcionários: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(funcionarios)
}

