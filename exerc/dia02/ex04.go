package main

import "fmt"

func main(){
	fmt.Println("Informe a sua idade: ")
	var idade int
	fmt.Scanln(&idade)
	if idade > 60 {
		fmt.Printf("A pessoa de %d anos é idoso!", idade)
	}else if idade >= 18 && idade<= 60{
		fmt.Printf("A pessoa de %d anos é maior de idade!", idade)
	}else if idade < 0{
		fmt.Println("O valor informado para idade é inválido!")
	}else {
		fmt.Printf("A pessoa de %d anos é menor de idade!", idade)
	}
}

