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
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/") // pega o primeiro argumento da cli
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()           // fecha a conexão no final
		res, err := io.ReadAll(req.Body) // lê a resposta com io.ReadAll
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}
		var data ViaCEP                  // cria uma variável do tipo ViaCEP
		err = json.Unmarshal(res, &data) // converte a resposta para ViaCEP
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter resposta: %v\n", err)
		}
		fmt.Println(data) // imprime a resposta

		file, err := os.Create("cidade.txt") // cria um arquivo
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()
		var address string = fmt.Sprintf("Cep: %s, Logradouro: %s, Complemento: %s, Bairro: %s, Localidade: %s, UF: %s, IBGE: %s, GIA: %s, DDD: %s, SIAFI: %s", data.Cep, data.Logradouro, data.Complemento, data.Bairro, data.Localidade, data.Uf, data.Ibge, data.Gia, data.Ddd, data.Siafi) // fecha o arquivo no final
		_, err = file.WriteString(address)                                                                                                                                                                                                                                                     // escreve no arquivo
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}
		fmt.Println("Arquivo criado com sucesso!")
		fmt.Println("Dados: ", address)
	}
}
