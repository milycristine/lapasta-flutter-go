package nota

import (
	"encoding/json"
	"lapasta/internal/models"
	"log"
	"net/http"
)

type NotaHandler struct {
	service NotaService
}

func NovaNotaHandler(service NotaService) *NotaHandler {
	return &NotaHandler{
		service: service,
	}
}

func (h *NotaHandler) CriarNota(w http.ResponseWriter, r *http.Request) {
	var nota models.Nota
	response := models.ResponseDefaultModel{
		IsSuccess: true,
	}

	if err := json.NewDecoder(r.Body).Decode(&nota); err != nil {
		log.Printf("Erro ao decodificar a nota: %v", err)
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Formato de entrada inválido"
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err := h.service.CriarNota(&nota); err != nil {
			log.Printf("Erro ao criar a nota: %v", err)
			response.IsSuccess = false
			response.Error = err
			response.ErrorMessage = "Erro ao criar a nota"
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			response.Data = nota
			w.WriteHeader(http.StatusCreated)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *NotaHandler) ListarNotas(w http.ResponseWriter, r *http.Request) {
	notas, err := h.service.ListarNotas()
	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      notas,
	}

	if err != nil {
		log.Printf("Erro ao listar notas: %v", err)
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao listar notas"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
