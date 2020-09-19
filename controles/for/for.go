package main

import (
	"fmt"
	"time"
)

func main() {
	n := 1
	for n <= 10 { // Não tem while em Go
		fmt.Println(n)
		n++
	}

	for x := 0; x <= 20; x += 2 {
		fmt.Println(x)
	}

	fmt.Println("\nMisturando...")
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	for {
		fmt.Println("Pra sempre")
		time.Sleep(time.Second)
	}
}
