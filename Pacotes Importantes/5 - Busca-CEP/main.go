package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		fmt.Println(os.Args[1:])

		// Relizando a chamada HTTP
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}

		// Sempre precisa fechar a requisição
		defer req.Body.Close()

		// Guardando a resposta em uma variável
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		// Trasnsformando o binário em JSON
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			panic(err)
		}

		// Criando um arquivo  
		file, err := os.Create("cidade.txt")
		if err != nil {
			panic(err)
		}

		// Também sempre precisamos fechar o arquivo
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s", data.Cep, data.Localidade))
	}
}