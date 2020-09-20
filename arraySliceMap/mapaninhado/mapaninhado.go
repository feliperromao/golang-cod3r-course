package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Guilherme": 1234.0,
			"Glauber":   4332.98,
		},
		"P": {
			"Paulo": 1212.2,
			"Pedro": 928.12,
		},
		"F": {
			"Felipe":    91010.0,
			"Fransicco": 28331.98,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Printf("--------- %s ---------\n", letra)
		for nome, salario := range funcs {
			fmt.Printf("%s R$ %.2f\n", nome, salario)
		}
	}
}
