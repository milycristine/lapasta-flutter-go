package recebimento

import (
	"encoding/json"
	utils "lapasta/internal/Utils"
	"lapasta/internal/models"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

func CriarRecebimento(w http.ResponseWriter, r *http.Request) {
	var recebimento models.Recebimento
	utils.SetSQLConn(nil)
	// Decodifica o JSON
	if err := json.NewDecoder(r.Body).Decode(&recebimento); err != nil {
		log.Println(err) // Adicione um log para verificar o erro
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insere
	if err := utils.ConnectionDb.CriarRecebimento(&recebimento); err != nil {
		log.Println(err) // Adicione um log para verificar o erro
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recebimento)
}

func ListarRecebimentos(w http.ResponseWriter, r *http.Request) {
	recebimentos, err := utils.ConnectionDb.ListarRecebimentos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna a lista de recebimentos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recebimentos)
}
