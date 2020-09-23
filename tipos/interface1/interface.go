package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	produto1 := produto{
		nome:  "Lapis",
		preco: 1.98,
	}
	pessoa1 := pessoa{
		nome:      "Jose",
		sobrenome: "Martins",
	}

	imprimir(produto1)
	imprimir(pessoa1)

	var coisa imprimivel = pessoa{"Fulano", "da Silva"}
	fmt.Println(coisa.toString())

	imprimir(produto{"Moleton", 49.90})
}
