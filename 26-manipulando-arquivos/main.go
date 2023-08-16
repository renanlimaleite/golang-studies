package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("1. Escrevendo no arquivo com Write / WriteString...")
	// tamanho, err := f.WriteString("Hello, World!")
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo..."))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes.\n", tamanho)
	f.Close()

	fmt.Println("2. Abrindo o arquivo e lendo com ReadFile")
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	fmt.Println("3. Leitura de pouco em pouco abrindo o arquivo com bufio")
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n])) // slice de 0 até n (10)
	}

	fmt.Println("4. Removendo o arquivo...")
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
