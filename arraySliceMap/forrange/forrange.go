package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // ... Compilador que conta
	for i, numero := range numeros {
		// fmt.Println(i, " - ", numero)
		fmt.Printf("%d) %d\n", i, numero)
	}

	// Ignorando o indice
	for _, num := range numeros {
		fmt.Println(num)
	}
}
