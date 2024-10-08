package funcionario

import (
	"encoding/json"
	"lapasta/internal/models"
	"net/http"
)

type FuncionarioHandler interface {
	CriarFuncionario(w http.ResponseWriter, r *http.Request)
	ListarFuncionarios(w http.ResponseWriter, r *http.Request)
}

type funcionarioHandler struct {
	service FuncionarioService
}

func NovoFuncionarioHandler(service FuncionarioService) FuncionarioHandler {
	return &funcionarioHandler{
		service: service,
	}
}

func (h *funcionarioHandler) CriarFuncionario(w http.ResponseWriter, r *http.Request) {
	var funcionario models.Funcionario
	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      funcionario,
	}

	if err := json.NewDecoder(r.Body).Decode(&funcionario); err != nil {
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao decodificar o funcionário"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.CriarFuncionario(&funcionario); err != nil {
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao criar funcionário"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *funcionarioHandler) ListarFuncionarios(w http.ResponseWriter, r *http.Request) {
	funcionarios, err := h.service.ListarFuncionarios()
	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      funcionarios,
	}

	if err != nil {
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao listar funcionários"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
