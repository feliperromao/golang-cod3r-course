package main

import "fmt"

func main() {
	p1 := Ponto{0.0, 0.0}
	p2 := Ponto{5.0, 5.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
