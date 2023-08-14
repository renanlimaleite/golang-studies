package main

// constraints
type MyNumber int

// ~ significa que considera que posso usar MyNumber como int ou float64
type Number interface {
	~int | float64
}

// generics
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// func Compara[T Number](a T, b T) bool {
// 	return a == b
// }

// comparable a > b não funciona porque não é comparável, exemplo: int e string
func Compara[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Renan": 1000, "João": 500, "Maria": 200}
	println(Soma(m)) // 1700

	m2 := map[string]float64{"Renan": 1000.50, "João": 500.50, "Maria": 200.50} // +1.701500e+003
	println(Soma(m2))

	m3 := map[string]MyNumber{"Renan": 1000, "João": 500, "Maria": 200}
	println(Soma(m3)) // 1700

	// println(Compara(10, 10.0)) // não consegue pois não é comparável int e float64
	println(Compara(10, 10)) // true
}
