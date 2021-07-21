package main

import "fmt"

func main(){
	fmt.Println("Digite um número:")
	var numero int
	fmt.Scanln(&numero)
	switch  {
	case numero == 0:
		fmt.Println("O número é 0.")
	case numero % 2 == 0:
		fmt.Println("O número é múltiplo de 2.")
	case numero % 3 == 0:
		fmt.Println("O número é múltiplo de 3.")
	default:
		fmt.Println("O número informado não é 0, nem múltiplo de 2 ou 3.")

	}


}
