package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais metodos
}

type bmwm7 struct{}

func (b bmwm7) ligarTurbo() {
	fmt.Println("Ligando turbo...")
}

func (b bmwm7) fazerBaliza() {
	fmt.Println("Fazendo baliza...")
}

func main() {
	var newBmwM7 esportivoLuxuoso = bmwm7{}
	newBmwM7.fazerBaliza()
	newBmwM7.ligarTurbo()
}
