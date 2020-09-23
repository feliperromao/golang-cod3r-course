package main

import "fmt"

type carro struct {
	nome            string
	velocidateAtual int
}

type ferrari struct {
	carro       // Campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
}
