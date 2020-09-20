package main

import "fmt"

func main() {
	// var aprovados map[int]string //! Map devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123] = "Felipe"
	aprovados[456] = "Marina"
	aprovados[789] = "Bia"
	aprovados[120] = "Guilherme"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	//? Lendo um valor de um map atraves da chave
	fmt.Println(aprovados[123])

	//? Excliindo um item de um map
	delete(aprovados, 123)
	fmt.Println(aprovados)
}
