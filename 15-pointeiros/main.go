package main

func main() {
	// Memória -> Endereço -> Valor
	a := 10
	println(&a) // 0xc00004e768
	var pointeiro *int = &a
	println(pointeiro) // 0xc00004e768
	*pointeiro = 20
	println(a) // 20 -> mudou valor na memória

	b := &a
	println(b)  // 0xc00004e768
	println(*b) // 20
	*b = 30     // mudou o valor na memoria para 30
	println(*b) // 30
	println(a)  // 30
}
