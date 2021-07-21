package main

import "fmt"

func main(){
	fmt.Println("Digite o primeiro valor:")
	var valor1 int
	fmt.Scanln(&valor1)
	fmt.Println("Digite o segundo valor:")
	var valor2 int
	fmt.Scanln(&valor2)
	fmt.Println("Digite o terceiro valor:")
	var valor3 int
	fmt.Scanln(&valor3)

	if valor1 > valor2 && valor1 > valor3{
		fmt.Println("O número de maior valor é: ", valor1)
	}else if valor2 > valor1 && valor2 > valor3{
		fmt.Println("O número de maior valor é: ", valor2)
	}else if valor3 > valor1 && valor3 > valor2{
		fmt.Println("O número de maior valor é: ", valor3)
	}


}
