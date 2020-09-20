package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aluno := range aprovados {
		fmt.Printf("%d) %s\n", i, aluno)
	}
}

func main() {
	alunos := []string{"Maria", "Pedro", "Felipe", "Guilherme"}
	imprimirAprovados(alunos...)
}
