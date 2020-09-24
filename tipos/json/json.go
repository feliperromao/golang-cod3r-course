package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Arroz", 4.89, []string{"mantimentos", "promocao"}}
	fmt.Println(p1)

	p1Json, err := json.Marshal(p1)
	if err == nil {
		fmt.Println(string(p1Json))
	}

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Feijao","preco":6.51,"tags":["mantimentos","escarco"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
