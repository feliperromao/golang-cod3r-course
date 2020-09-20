package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // Reteorno limpo
}

func main() {
	n1, n2 := trocar(10, 20)
	fmt.Println(n1, n2)
}
