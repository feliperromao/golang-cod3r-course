package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // switch true, vai executar o primeiro case que a expressão seja true
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}
}
