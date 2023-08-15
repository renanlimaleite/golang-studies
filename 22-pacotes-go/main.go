package main

import (
	"curso-go-teste/matematica"
	"fmt"
)

func main() {
	soma := matematica.Soma(10, 20)

	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Printf("O resultado é %v\n", soma)

	fmt.Println(matematica.A)

	fmt.Println(carro)

	fmt.Println(carro.Andar()) // funciona pois é público letra maiúscula
	// fmt.Println(carro.andar()) -> não funciona pois é privado letra minúscula
}
