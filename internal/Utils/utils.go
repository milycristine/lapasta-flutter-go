package utils

import (
	_ "embed" // EMB
	sql "lapasta/database"
	"log"
)

// ConnectionDb ..
var ConnectionDb *sql.SQLStr

// SetSQLConn armazena a conexão com o banco de dados.
func SetSQLConn(l *sql.SQLStr) {
	if l == nil {
		log.Fatal("A conexão com o banco de dados não pode ser nula.")
	}
	ConnectionDb = l
}
