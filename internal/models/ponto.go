package models

type Ponto struct {
    Id             int    `json:"id"`
    HManha         string `json:"hManha"`
    HAlmocoEntrada string `json:"hAlmocoEntrada"`
    HAlmocoSaida   string `json:"hAlmocoSaida"`
    HNoite         string `json:"hNoite"`
    Dia            string `json:"dia"`
    Situacao       string `json:"situacao"`
    IdFuncionario  int    `json:"idFuncionario"`
}
