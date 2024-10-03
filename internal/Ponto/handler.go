// handler/ponto.go
package ponto

import (
	"encoding/json"
	"lapasta/internal/models"
	"log"
	"net/http"
)

type PontoHandler struct {
	service PontoService
	
}

func NovoPontoHandler(service PontoService) *PontoHandler {
	return &PontoHandler{
		service: service,
	}
}

func (h *PontoHandler) CriarPonto(w http.ResponseWriter, r *http.Request) {
	var ponto models.Ponto

	if err := json.NewDecoder(r.Body).Decode(&ponto); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CriarPonto(&ponto); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ponto)
}

func (h *PontoHandler) ListarPontos(w http.ResponseWriter, r *http.Request) {
	pontos, err := h.service.ListarPontos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pontos)
}
