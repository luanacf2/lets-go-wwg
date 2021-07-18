package main

import "fmt"

func main() {
	listaDeMercado := []string{"abacate", "sabonete", "azeite", "tomate", "banana"}
	for i, v := range listaDeMercado{
		fmt.Printf("%d) %s \n", i+1, v)
	}
}