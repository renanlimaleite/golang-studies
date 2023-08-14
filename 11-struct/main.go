package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	renan := Cliente{
		Nome:  "Renan",
		Idade: 26,
		Ativo: true,
	}

	renan.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", renan.Nome, renan.Idade, renan.Ativo)
}
