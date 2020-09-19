package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	
	p = &i // Pegando o endereço da variavel
	*p++   // Desreferenciando (pegando o valor)
	i++
	
	// Go nao tem aritmetica de ponteiros
	// p++

	fmt.Println(p, *p, i)
}
