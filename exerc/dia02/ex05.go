package main

import "fmt"

func main()  {
	fmt.Println("Informe a sua idade: ")
	var idade int
	fmt.Scanln(&idade)
	switch {
	case idade > 60:
		fmt.Printf("A pessoa de %d anos é idoso!", idade)
	case idade >= 18 && idade<= 60:
		fmt.Printf("A pessoa de %d anos é maior de idade!", idade)
	case idade < 0:
		fmt.Println("O valor informado para idade é inválido!")
	default:
		fmt.Printf("A pessoa de %d anos é menor de idade!", idade)
	}

}
