package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Name   string `json:"nm"`
	Numero int    `json:"n"`
	Saldo  int    `json:"s"`
}

func main() {
	conta := Conta{Name: "Renan", Numero: 123, Saldo: 1000}

	// como transformar a struct em json?
	fmt.Println("1. Transformando a struct em json...")
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// criando um json com NewEncoder -> Entregando um json pronto no caso, para o os.Stdout
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	/*
		nesse caso tem que ser igual ao json que está sendo passado
	*/
	// jsonPuro := []byte(`{"Name":"Renan","Numero":2,"Saldo":2000}`)
	// var contaX Conta
	// // transformando o json em struct
	// fmt.Println("2. Transformando o json em struct...")
	// err = json.Unmarshal(jsonPuro, &contaX)
	// if err != nil {
	// 	panic(err)
	// }
	// println(contaX.Name, contaX.Numero, contaX.Saldo)

	/*
		nesse caso funciona pq a struct tem bind com o json
	*/
	jsonPuro2 := []byte(`{"nm":"Renan","n":2,"s":2000}`)
	var contaY Conta
	// transformando o json em struct
	fmt.Println("2. Transformando o json em struct...")
	err = json.Unmarshal(jsonPuro2, &contaY)
	if err != nil {
		panic(err)
	}
	println(contaY.Name, contaY.Numero, contaY.Saldo)
}
