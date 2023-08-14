package main

import "fmt"

type Cliente struct {
	nome string
}

// atribui a uma struct (c Cliente) um método "andou()"
func (c Cliente) andou() {
	c.nome = "Wesley Willians"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

// func main() {
// 	// cria uma instância do struct Cliente com o nome 'Wesley'"
// 	wesley := Cliente{
// 		nome: "Wesley",
// 	}
// 	// automaticamente essa variavel wesley tem acesso ao método andou()
// 	wesley.andou() // O cliente Wesley Willians andou

// 	// O valor da struct com nome Wesley
// 	fmt.Printf("O valor da struct com nome %v\n", wesley.nome)
// }
