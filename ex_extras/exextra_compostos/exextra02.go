package main

import "fmt"

func main(){
	vermelho := []string{"Helena", "Jonas", "José", "Juliana"}
	vermelho = append(vermelho, "Luis Inácio")

	fmt.Println("Time vermelho: ",vermelho)
}
