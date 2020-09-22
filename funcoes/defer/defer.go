package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	obterValorAprovado(9000)
	obterValorAprovado(2000)
}
