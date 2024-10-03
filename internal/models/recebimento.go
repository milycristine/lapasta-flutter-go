package models

type Recebimento struct {
	Id            int    `json:"id"`
	Dia           string `json:"dia"`
	Empresa       string `json:"empresa"`
	Produto       string `json:"produto"`
	UrlImagem     string `json:"urlImagem"`
	IdResponsavel uint   `json:"idResponsavel"`
}
