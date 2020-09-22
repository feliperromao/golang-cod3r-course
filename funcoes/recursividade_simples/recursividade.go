package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n - 1)
	}
}

func main() { 
	fmt.Println(fatorial(5))
	fmt.Println(fatorial(9))
	fmt.Println(fatorial(0))
}