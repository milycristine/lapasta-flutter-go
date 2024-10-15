package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config representa a estrutura do arquivo de configuração.
type Config struct {
	API struct {
		Host string `yaml:"host,omitempty"`
		Port string `yaml:"port,omitempty"`
	} `yaml:"api,omitempty"`
	SQL struct {
		Host     string `yaml:"host,omitempty"`
		Port     string `yaml:"port,omitempty"`
		User     string `yaml:"username,omitempty"`
		Password string `yaml:"password,omitempty"`
	} `yaml:"sql,omitempty"`
}

var Yml Config

func LoadConfig() error {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &Yml)
}

func CreateConfigFile() {
	if _, err := os.Stat("config.yaml"); err == nil {
		fmt.Println("O arquivo 'config.yaml' já existe. Deseja sobrescrever? (y/N)")
		var rsp string
		fmt.Scan(&rsp)
		if strings.ToLower(rsp) == "y" {
			writeFile()
		}
		return
	}
	writeFile()
}

func writeFile() {
	data, err := yaml.Marshal(Yml)
	if err != nil {
		fmt.Printf("Erro ao gerar o YAML: %v\n", err)
		return
	}
	if err := os.WriteFile("config.yaml", data, 0644); err != nil {
		fmt.Printf("Erro ao escrever no arquivo config.yaml: %v\n", err)
	}
}
