package auth

import (
	"encoding/json"
	"lapasta/internal/models"
	"net/http"
)

func LoginHandler(s AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginRequest models.Login
		response := models.ResponseDefaultModel{
			IsSuccess: true,
		}

		if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
			response.IsSuccess = false
			response.Error = err
			response.ErrorMessage = "Erro ao decodificar a requisição de login"
			w.WriteHeader(http.StatusBadRequest)
		} else {
			login, err := s.Login(loginRequest.Username, loginRequest.Password)
			if err != nil {
				response.IsSuccess = false
				response.Error = err
				response.ErrorMessage = "Usuário ou senha inválidas"
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				response.Data = login
				w.WriteHeader(http.StatusOK)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
