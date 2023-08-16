// defer é "atrasar" alguma coisa para ser executada por último
// útil para fechar conexões com banco de dados, arquivos, etc

package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("1. Primeira Linha")
	fmt.Println("2. Segunda Linha")
	defer fmt.Println("3. Terceira Linha")
}
