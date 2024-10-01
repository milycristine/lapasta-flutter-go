package server

import (
    "embed"
    "log"
    "net/http"
    "lapasta/api"
    config "lapasta/config"
)

var fs embed.FS

// Controllers inicia o servidor HTTP e define a rota de login.
func Controllers() {
    log.Printf("Iniciando servidor na porta: %s", config.Yml.API.Port)

    http.HandleFunc("/login", api.Login)
    http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.FS(fs))))
    
    log.Fatal(http.ListenAndServe(":"+config.Yml.API.Port, nil))
}
