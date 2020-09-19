package main

import "fmt"

func notaParaConceito(nota float64) string {
	switch {
	case nota <= 10 && nota >= 9:
		return "A"
	case nota < 9 && nota >= 8:
		return "B"
	case nota < 8 && nota >= 6:
		return "C"
	case nota < 6 && nota >= 3:
		return "D"
	case nota < 3 && nota >= 0:
		return "E"
	default:
		return "Nota Invalida"
	}
}

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(11))
	fmt.Println(notaParaConceito(-1))
	fmt.Println(notaParaConceito(4))
	fmt.Println(notaParaConceito(1))
	fmt.Println(notaParaConceito(7.65))
	fmt.Println(notaParaConceito(8.2))
}
