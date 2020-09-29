package main

import (
	"fmt"
	"time"
)

func isPrimo(numero int) bool {
	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Second)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c)

	for i := range c {
		fmt.Printf("%d ", i)
	}
	fmt.Println("Fim")
}
