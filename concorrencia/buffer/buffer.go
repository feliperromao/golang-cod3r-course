package main

import "fmt"

func rotina(c chan int) {
	fmt.Println("Iniciou")
	c <- 1
	fmt.Println("Iniciou - 1")
	c <- 2
	fmt.Println("Iniciou - 2")
	c <- 3
	fmt.Println("Iniciou - 3")
	c <- 4
	fmt.Println("Iniciou - 4")
	c <- 5
	fmt.Println("Iniciou - 5")
	c <- 6
	fmt.Println("Iniciou - 6")
}

func main() {
	c := make(chan int, 3)

	go rotina(c)

	fmt.Println(<-c)
}
