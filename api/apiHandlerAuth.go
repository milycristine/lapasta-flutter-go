package api

import (
	"encoding/json"
	"fmt"
	"lapasta/models"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

// Requisição de login
func Login(w http.ResponseWriter, r *http.Request) {
	client := models.Login{}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		sendSimpleResponse(w, "ERROR", "Bad Request", http.StatusBadRequest)
		return
	}

	if err := validateCredentials(w, client.Username, client.Password); err != nil {
		sendSimpleResponse(w, "ERROR", err.Error(), http.StatusUnauthorized)
		return
	}
	sendSuccessResponse(w, client.Username)
}

// Função para validar as credenciais do usuário
func validateCredentials(w http.ResponseWriter, username, senha string) error {
	log.Printf("Validando usuário: %s", username)

	data, err := connectionDb.Autenticar(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status: "Bad Request",
			Error:  "",
			Data:   err.Error(),
		})

	}

	// Compara a senha fornecida com o hash armazenado
	if !comparePasswords(data.Password, senha) {
		return fmt.Errorf("Nome ou senha inválidos")
	}

	return nil
}

// Função para comparar a senha fornecida com o hash armazenado
func comparePasswords(hash string, senha string) bool {
	if hash == senha {
		return true
	}
	return false
}

// Funções de resposta simples e sucesso
func sendSimpleResponse(w http.ResponseWriter, status, errorMsg string, code int) {
	response := map[string]interface{}{
		"status": status,
		"error":  errorMsg,
	}

	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Erro ao codificar resposta simples: %v", err)
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
