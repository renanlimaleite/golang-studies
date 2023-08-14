package main

import "fmt"

func main() {
	var minhaVar interface{} = "Renan Leite"

	println(minhaVar)          // (0x45d6a0,0x47c718)
	println(minhaVar.(string)) // Renan Leite
	res, ok := minhaVar.(int)

	// O valor de res é 0 e o valor de ok é false
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	res2 := minhaVar.(int)
	fmt.Println(res2) // panic: interface conversion: interface {} is string, not int
}
