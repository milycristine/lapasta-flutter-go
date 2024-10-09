package recebimento

import (
	"encoding/json"
	"lapasta/internal/models"
	"log"
	"net/http"
)

type RecebimentoHandler struct {
	service RecebimentoService
}

func NovoRecebimentoHandler(service RecebimentoService) *RecebimentoHandler {
	return &RecebimentoHandler{
		service: service,
	}
}

func (h *RecebimentoHandler) CriarRecebimento(w http.ResponseWriter, r *http.Request) {
	var recebimento models.Recebimento
	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      recebimento,
	}

	if err := json.NewDecoder(r.Body).Decode(&recebimento); err != nil {
		log.Println(err)
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao decodificar o recebimento"
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := h.service.CriarRecebimento(&recebimento); err != nil {
		log.Println(err)
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao criar o recebimento"
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = recebimento
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *RecebimentoHandler) ListarRecebimentos(w http.ResponseWriter, r *http.Request) {
	recebimentos, err := h.service.ListarRecebimentos()
	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      recebimentos,
	}

	if err != nil {
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao listar recebimentos"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
