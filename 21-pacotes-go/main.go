package main

import (
	"fmt"

	"curso-go/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)

	fmt.Printf("O resultado é %v\n", soma)
}
