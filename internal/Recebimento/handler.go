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

	// Decodifica o JSON
	if err := json.NewDecoder(r.Body).Decode(&recebimento); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insere
	if err := h.service.CriarRecebimento(&recebimento); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recebimento)
}

func (h *RecebimentoHandler) ListarRecebimentos(w http.ResponseWriter, r *http.Request) {
	recebimentos, err := h.service.ListarRecebimentos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna a lista de recebimentos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recebimentos)
}
