package main

import (
	"fmt"
	"time"
)

func fale(pessoa, frase string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, frase, i+1)
	}
}

func main() {
	// fale("Maria", "pq vc nao fala comigo?", 2)
	// fale("Joao", "So posso depois de vc falar", 1)

	// go fale("Maria", "Eeei...", 3)
	// go fale("Joao", "Opa...", 2)

	go fale("Maria", "Entendi!!!", 10)
	fale("Jose", "Parabens!!!", 5)
}
