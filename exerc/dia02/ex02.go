package main

import (
	"fmt"
	"time"
)

func main(){
	//Escreva um programa que verifique se no ano atual essa pessoa já está
	//apta a votar e que printe na tela uma mensagem informando caso positivo e caso negativo.
	nasci := 2001
	ano := time.Now().Year()
	idade := ano - nasci

	if idade >= 16{
		fmt.Printf("Está pessoa tem %d anos de idade e está apto para votar!", idade)
		return
	}
	fmt.Printf("Está pessoa tem %d anos de idade e não está apto para votar!", idade)
}
