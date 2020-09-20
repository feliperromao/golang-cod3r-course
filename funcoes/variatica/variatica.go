package main

import "fmt"

func media(numeros ...float64) float64 {
	count := len(numeros)
	if count == 0 {
		return 0.0
	}

	total := 0.
	for _, num := range numeros {
		total += num
	}

	return total / float64(count)
}

func main() {
	fmt.Println(media(1,2))
	fmt.Println(media(1,2,3,4,5,6,7,8,9,10))
	fmt.Println(media())
}
