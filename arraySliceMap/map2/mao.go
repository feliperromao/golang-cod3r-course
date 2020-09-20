package main

import "fmt"

func main() {
	salarios := map[string]float64{
		"Jose Martins":   8945.19,
		"Maria da Penha": 12984.00,
		"Pedro Costa":    865.19,
	}
	salarios["Marina Lacerda"] = 3900.45
	fmt.Println(salarios)

	//? Nao gera erro ao tentar excluir um item que nao existe
	delete(salarios, "inexistente")

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}
}
