package main

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println("Simulação ->", c.saldo) // Simulação -> 300
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)

	// 300 -> mudou para 300 porque a função simular está recebendo
	// o endereço da conta 'c' na memória e alterando o valor de c.saldo na memória
	// (teoria de pointeiros)
	println(conta.saldo)
}
