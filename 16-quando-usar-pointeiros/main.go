package main

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	// println(soma(minhaVar1, minhaVar2)) // 30
	soma(&minhaVar1, &minhaVar2)
	// era 10, mas mandei o endereço, na função mudou na memoria para 50 então agora é 50
	println(minhaVar1) // 50
	println(minhaVar2) // 50
}
