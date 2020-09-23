package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual uint
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	f40 := ferrari{
		modelo:          "F40",
		turboLigado:     false,
		velocidadeAtual: 0,
	}
	f40.ligarTurbo()
	fmt.Println(f40)

	var spider esportivo = &ferrari{"Spider", false, 0}
	spider.ligarTurbo()
	fmt.Println(spider)
}
