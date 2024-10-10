package models

type Nota struct {
	Id            int     `json:"id"`
	Tipo          string  `json:"tipo"`
	Valor         float64 `json:"valor"`
	IdFuncionario int     `json:"idFuncionario"`
	UrlImagem     string  `json:"urlImagem"`
	Dia           string  `json:"dia"`
	NomeFuncionario string `json:"nomeFuncionario"` 

}
