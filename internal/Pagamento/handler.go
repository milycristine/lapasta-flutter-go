package pagamento

import (
	"encoding/json"
	"lapasta/internal/models"
	"log"
	"net/http"
	"strconv"
)

type PagamentoHandler struct {
	service PagamentoService
}

func NovoPagamentoHandler(service PagamentoService) *PagamentoHandler {
	return &PagamentoHandler{
		service: service,
	}
}

func (h *PagamentoHandler) CriarPagamento(w http.ResponseWriter, r *http.Request) {
	var pagamento models.Pagamento
	response := models.ResponseDefaultModel{
		IsSuccess: true,
	}

	if err := json.NewDecoder(r.Body).Decode(&pagamento); err != nil {
		log.Printf("Erro ao decodificar o pagamento: %v", err)
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Formato de entrada inválido"
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err := h.service.CriarPagamento(&pagamento); err != nil {
			log.Printf("Erro ao criar o pagamento: %v", err)
			response.IsSuccess = false
			response.Error = err
			response.ErrorMessage = "Erro ao criar o pagamento"
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			response.Data = pagamento
			w.WriteHeader(http.StatusCreated)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PagamentoHandler) ListarPagamentos(w http.ResponseWriter, r *http.Request) {
	pagamentos, err := h.service.ListarPagamentos()
	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      pagamentos,
	}

	if err != nil {
		log.Printf("Erro ao listar pagamentos: %v", err)
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao listar pagamentos"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (h *PagamentoHandler) ListarPagamentosPorDia(w http.ResponseWriter, r *http.Request) {
	diaStr := r.URL.Query().Get("dia")
	dia, err := strconv.Atoi(diaStr)
	if err != nil || (dia != 5 && dia != 20) {
		http.Error(w, "Parâmetro dia deve ser 5 ou 20", http.StatusBadRequest)
		return
	}

	pagamentos, err := h.service.ListarPagamentosPorDia(dia)
	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      pagamentos,
	}

	if err != nil {
		log.Printf("Erro ao listar pagamentos: %v", err)
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao listar pagamentos"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (h *PagamentoHandler) AtualizarPagamento(w http.ResponseWriter, r *http.Request) {
	var pagamento models.Pagamento
	response := models.ResponseDefaultModel{
		IsSuccess: true,
	}

	if err := json.NewDecoder(r.Body).Decode(&pagamento); err != nil {
		log.Printf("Erro ao decodificar o pagamento: %v", err)
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Formato de entrada inválido"
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err := h.service.AtualizarStatusPagamento(pagamento.Id, pagamento.StatusId); err != nil {
			log.Printf("Erro ao atualizar o pagamento: %v", err)
			response.IsSuccess = false
			response.Error = err
			response.ErrorMessage = "Erro ao atualizar o pagamento"
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNoContent) 
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}