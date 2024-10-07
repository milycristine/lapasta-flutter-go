package server

import (
	"embed"
	"log"
	"net/http"

	config "lapasta/config"
	auth "lapasta/internal/AUTH"
	funcionario "lapasta/internal/Funcionario"
	ponto "lapasta/internal/Ponto"
	recebimento "lapasta/internal/Recebimento"
	utils "lapasta/internal/Utils"
)

var fs embed.FS

func Controllers() {
	repo := auth.NewAuthRepository(utils.ConnectionDb)
	authService := auth.NewAuthService(repo)

	recebimentoRepo := recebimento.NovoRecebimentoRepository(utils.ConnectionDb) 
	recebimentoService := recebimento.NovoRecebimentoService(recebimentoRepo)
	recebimentoHandler := recebimento.NovoRecebimentoHandler(recebimentoService)

	pontoRepo := ponto.NovoPontoRepository(utils.ConnectionDb)
	pontoService := ponto.NovoPontoService(pontoRepo)
	pontoHandler := ponto.NovoPontoHandler(pontoService)

	funcionarioRepo := funcionario.NovoFuncionarioRepository(utils.ConnectionDb)
	funcionarioService := funcionario.NovoFuncionarioService(funcionarioRepo)
	funcionarioHandler := funcionario.NovoFuncionarioHandler(funcionarioService)

	log.Printf("Iniciando servidor na porta: %s", config.Yml.API.Port)

	http.HandleFunc("/login", auth.LoginHandler(authService))

	http.HandleFunc("/recebimento", recebimentoHandler.CriarRecebimento)         
	http.HandleFunc("/listarRecebimento", recebimentoHandler.ListarRecebimentos) 

	http.HandleFunc("/baterPonto", pontoHandler.CriarPonto)   
	http.HandleFunc("/listarPonto", pontoHandler.ListarPontos) 

	http.HandleFunc("/funcionario", funcionarioHandler.CriarFuncionario)         
	http.HandleFunc("/listarFuncionario", funcionarioHandler.ListarFuncionarios) 

	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.FS(fs))))

	log.Fatal(http.ListenAndServe(":"+config.Yml.API.Port, nil))
}
