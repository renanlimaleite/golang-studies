package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Wesley"]) // 1000
	delete(salarios, "Wesley")      // deleta da chave
	fmt.Println(salarios["Wesley"]) // 0
	salarios["Wes"] = 5000
	fmt.Println(salarios["Wes"]) // 5000

	sal := make(map[string]int) // criando map do zero
	sal["Renan"] = 20000
	sal1 := map[string]int{}
	sal1["Wesley"] = 1000

	fmt.Println(sal["Renan"])   // 20000
	fmt.Println(sal1["Wesley"]) // 1000

	// percorrendo map
	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é de %d\n", nome, salario)
	}

	// percorrendo map ignorando nome _ = blank indetify
	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
