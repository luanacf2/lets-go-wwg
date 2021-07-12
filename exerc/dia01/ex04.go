package main

import (
	"fmt"
)

func main() {
	nome:= "Luana"
	cor := "Preto"

	fmt.Printf("Meu nome é %s e minha cor preferida é %s.\n", nome, cor)

	cor = "Azul"

	fmt.Printf("Meu nome é %s e minha cor preferida é %s.", nome, cor)

}
