package main

import "fmt"

func main() {
	slice1 := make([]int, 10)
	slice1[9] = 12
	fmt.Println(slice1)

	slice1 =  make([]int, 10, 20) // 20 especifica o tamanho do array que esta sendo referencia do pelo slice
	fmt.Println(slice1, len(slice1), cap(slice1)) // cap mostra a capacidade maxima do slice: 20

	slice1 = append(slice1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 1)
	fmt.Println(slice1, len(slice1), cap(slice1))
}
