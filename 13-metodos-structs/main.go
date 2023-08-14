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
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado: %t", c.Nome, c.Ativo)
}

func main() {
	renan := Cliente{
		Nome:  "Renan",
		Idade: 26,
		Ativo: true,
	}

	renan.Desativar()
}
