package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

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

/*
QUALQUER FUNC QUE TIVER O MÉTODO DESATIVAR PODERÁ USAR O DESATIVACAO
*/
// INTERFACE SOMENTE MÉTODOS
type Pessoa interface {
	Desativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	renan := Cliente{
		Nome:  "Renan",
		Idade: 26,
		Ativo: true,
	}
	renan.Desativar()
	Desativacao(renan)

	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)
}
