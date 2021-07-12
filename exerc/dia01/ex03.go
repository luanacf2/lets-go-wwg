package main

import (
	"fmt"
)

func main() {
	var unidadesBrocolis, unidadesMacarrao, pesoDeMacas float64
	unidadesBrocolis = 3
	unidadesMacarrao = 2
	pesoDeMacas = 0.900

	var precoBrocolis, precoMacarrao, precoDeMacas float64
	precoBrocolis = 3.99
	precoMacarrao = 5.49
	precoDeMacas = 3.50

	valorDaCompra := (unidadesBrocolis * precoBrocolis) + (unidadesMacarrao * precoMacarrao) + (pesoDeMacas * precoDeMacas)

	fmt.Printf("O valor da compra deu: \n%v \n\t e o valor do Peso das Maças é de %v", valorDaCompra, precoDeMacas)


}