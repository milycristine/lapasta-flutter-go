package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"lapasta/models"
)

//requisição de login
func Login(w http.ResponseWriter, r *http.Request) {
	client := models.Login{}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		sendSimpleResponse(w, "ERROR", "Bad Request", http.StatusBadRequest)
		return
	}

	if err := validateCredentials(client.Email, client.Senha); err != nil {
		sendSimpleResponse(w, "ERROR", err.Error(), http.StatusUnauthorized)
		return
	}
	sendSuccessResponse(w, client.Email)
}

func validateCredentials(email, senha string) error {
	log.Printf("Validando usuário: %s", email)
	if email != "usuario@exemplo.com" || senha != "senha123" {
		return fmt.Errorf("Nome ou senha inválidos")
	}
	return nil
}

func sendSimpleResponse(w http.ResponseWriter, status, errorMsg string, code int) {
	response := map[string]interface{}{
		"status": status,
		"error":  errorMsg,
	}

	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Erro ao codificar resposta simples: %v", err) // Log do erro
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func sendSuccessResponse(w http.ResponseWriter, email string) {
	response := map[string]interface{}{
		"status":  "SUCCESS",
		"message": "Login bem-sucedido",
		"email":   email,
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Erro ao codificar resposta de sucesso: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
