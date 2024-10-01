package api

import (
	_ "embed" // EMB
	sql "lapasta/database"
	"log"
)

var connectionLinx *sql.SQLStr

// SetSQLConn armazena a conexão com o banco de dados.
func SetSQLConn(l *sql.SQLStr) {
	if l == nil {
		log.Fatal("A conexão com o banco de dados não pode ser nula.")
	}
	connectionLinx = l
}
