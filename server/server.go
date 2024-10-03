package server

import (
	"embed"
	config "lapasta/config"
	auth "lapasta/internal/AUTH"
	recebimento "lapasta/internal/Recebimento"
	utils "lapasta/internal/Utils"

	"log"
	"net/http"
)

var fs embed.FS

// Controllers inicia o servidor HTTP e define a rota de login.
func Controllers() {
	repo := auth.NewAuthRepository(utils.ConnectionDb) // Criar uma nova instância do repositório
	authService := auth.NewAuthService(repo)           // Criar uma nova instância do serviço de autenticação

	log.Printf("Iniciando servidor na porta: %s", config.Yml.API.Port)

	http.HandleFunc("/login", auth.LoginHandler(authService))
	http.HandleFunc("/recebimento", recebimento.CriarRecebimento)
	http.HandleFunc("/listarRecebimento", recebimento.ListarRecebimentos)
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.FS(fs))))

	log.Fatal(http.ListenAndServe(":"+config.Yml.API.Port, nil))
}
