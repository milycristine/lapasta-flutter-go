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
	response := models.ResponseDefaultModel{
		IsSuccess: true,
	}

	if err := json.NewDecoder(r.Body).Decode(&ponto); err != nil {
		log.Printf("Erro ao decodificar o ponto: %v", err)
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Formato de entrada inválido"
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err := h.service.CriarPonto(&ponto); err != nil {
			log.Printf("Erro ao criar o ponto: %v", err)
			response.IsSuccess = false
			response.Error = err
			response.ErrorMessage = "Erro ao criar ponto"
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			response.Data = ponto
			w.WriteHeader(http.StatusCreated)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PontoHandler) ListarPontos(w http.ResponseWriter, r *http.Request) {
	pontos, err := h.service.ListarPontos()
	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      pontos,
	}

	if err != nil {
		log.Printf("Erro ao listar pontos: %v", err)
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao listar pontos"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
