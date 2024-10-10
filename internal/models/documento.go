package models

import "time"

type Documento struct {
	Id            int       `json:"id"`
	Titulo        string    `json:"titulo"`
	Url           string    `json:"url"`
	DataCriacao   time.Time `json:"dataCriacao"`
	IdFuncionario int       `json:"idFuncionario"`
}
