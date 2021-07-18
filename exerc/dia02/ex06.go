package main

import "fmt"

func main(){
	fmt.Println("Informe a hora: ")
	var hora int
	fmt.Scanln(&hora)
	switch  {
	case hora >= 6 && hora < 12:
		fmt.Println("Está de manhã!")
	case hora >= 12 && hora < 18:
		fmt.Println("Está de tarde")
	case hora >= 18 && hora < 24:
		fmt.Println("Está de  noite")
	case hora >= 0 && hora <6:
		fmt.Println("Está de madrugada")
	default:
		fmt.Println("O valor informado não é uma hora válida!!")
	}
}
