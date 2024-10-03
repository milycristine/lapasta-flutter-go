// internal/auth/handler.go
package auth

import (
	"encoding/json"
	"lapasta/internal/models"
	"net/http"
)

// LoginHandler gerencia a requisição de login.
func LoginHandler(s AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginRequest models.Login
		if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		login, err := s.Login(loginRequest.Username, loginRequest.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Retorna a resposta de sucesso
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(login)
	}
}
