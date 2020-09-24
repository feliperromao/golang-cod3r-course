package main

import "fmt"

type curso struct {
	name string
}

func main() {
	var coisa interface{}
	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = true
	fmt.Println(coisa2)
	coisa2 = "Teste"
	fmt.Println(coisa2)
	coisa2 = curso{"Jose Felipe"}
	fmt.Println(coisa2)
}
