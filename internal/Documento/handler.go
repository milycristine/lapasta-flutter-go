package documento

import (
	"encoding/json"
	"lapasta/internal/models"
	"net/http"
)

type DocumentoHandler interface {
	CriarDocumento(w http.ResponseWriter, r *http.Request)
	ListarDocumentos(w http.ResponseWriter, r *http.Request)
}

type documentoHandler struct {
	service DocumentoService
}

func NovoDocumentoHandler(service DocumentoService) DocumentoHandler {
	return &documentoHandler{
		service: service,
	}
}

func (h *documentoHandler) CriarDocumento(w http.ResponseWriter, r *http.Request) {
	var documento models.Documento
	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      documento,
	}

	if err := json.NewDecoder(r.Body).Decode(&documento); err != nil {
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao decodificar o documento"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.CriarDocumento(&documento); err != nil {
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao criar documento"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *documentoHandler) ListarDocumentos(w http.ResponseWriter, r *http.Request) {
	documentos, err := h.service.ListarDocumentos()
	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      documentos,
	}

	if err != nil {
		response.IsSuccess = false
		response.Error = err
		response.ErrorMessage = "Erro ao listar documentos"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
