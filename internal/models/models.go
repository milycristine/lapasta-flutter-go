package models

// Funcionario representa os dados do funcionário no banco de dados.
type Funcionario struct {
	Id           int     `json:"id,omitempty"`
	Nome         string  `json:"nome,omitempty"`
	Sobrenome    string  `json:"sobrenome,omitempty"`
	Cpf          string  `json:"cpf,omitempty"`
	Rg           string  `json:"rg,omitempty"`
	DataNasc     string  `json:"data_nasc,omitempty"`
	Email        string  `json:"email,omitempty"`
	Senha        []byte  `json:"senha,omitempty"` // Senha em formato hash
	Cargo        string  `json:"cargo,omitempty"`
	DateAdmissao string  `json:"data_admissao,omitempty"`
	DateFinal    *string `json:"data_final,omitempty"`
	HoraEntrada  string  `json:"hora_entrada,omitempty"`
	HoraSaida    string  `json:"hora_saida,omitempty"`
}

type Login struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

//Response ...
type Response struct {
	Status string
	Error  string
	Data   interface{}
}
