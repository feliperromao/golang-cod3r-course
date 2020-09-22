package main

import "fmt"

type item struct {
	produtoID  uint
	quantidade uint
	preco      float64
}

type pedido struct {
	userID uint
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total
}

func main() {
	pedido1 := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, quantidade: 20, preco: 19.2},
			item{produtoID: 2, quantidade: 10, preco: 3.0},
			item{produtoID: 3, quantidade: 15, preco: 4.5},
		},
	}

	fmt.Println(pedido1.valorTotal())
}
