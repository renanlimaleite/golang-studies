package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	// Address Endereco
	Endereco
}

func main() {
	renan := Cliente{
		Nome:  "Renan",
		Idade: 26,
		Ativo: true,
	}

	renan.Ativo = false

	renan.Cidade = "São Paulo"                 // short
	renan.Endereco.Logradouro = "Rua Leopoldo" // long

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", renan.Nome, renan.Idade, renan.Ativo)
}
