package main

import "fmt"

func main(){
	listaDeMercado := []string{"abacate", "sabonete", "azeite de oliva", "tomate", "banana", "macarrão", "cebola"}

	for lista := 0; lista < 7; lista++{
		fmt.Printf("%d) %s\n", lista + 1 , listaDeMercado[lista])
	}
}