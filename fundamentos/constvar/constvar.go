package main

import (
	out "fmt" // Podemos alterar o nome no import do pacote
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // O tipo (float64) foi inferido pelo compilador

	// Forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)

	out.Println("Area da circuferencia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	out.Println(a, b, c, d)

	// Criando varias variaveis tipadas e atribuindo valores
	var e, f bool = true, false
	out.Println(e, f)

	g, h, i := 2, false, "epa!"
	out.Println(g, h, i)
}
