package main

import (
	"fmt"
	"reflect"
)

func main() {
	array1 := [...]int{1, 2, 3} // Array
	fmt.Println("Array", array1)
	fmt.Println(reflect.TypeOf(array1))

	slice1 := []int{1, 2, 3} //slice
	fmt.Println("Slice", slice1)
	fmt.Println(reflect.TypeOf(slice1))

	array2 := [...]int{1, 2, 3, 4, 5}

	// Slice não é um array
	// Slice define um pedaço de um array
	slice2 := array2[1:3]
	fmt.Println(slice2)

	slice3 := array2[:2] // Novo slice apontando para o mesmo array
	fmt.Println(slice3)

	// Slice a partir e um slice
	slice4 := slice3[:1]
	fmt.Println(slice4)
}
