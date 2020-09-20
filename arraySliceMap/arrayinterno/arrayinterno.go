package main

import "fmt"

func main() {
	array1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	slice1 := array1[:]
	slice2 := append(slice1, 123)

	fmt.Println(slice1, slice2)
	slice1[2] = 5
	fmt.Println(slice1, slice2)
}
