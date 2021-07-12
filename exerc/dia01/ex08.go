package main

import "fmt"

func main(){
	ano := make(map[int]string)
	ano[1] = "Janeiro"
	ano[2] = "Fevereiro"
	ano[3] = "Março"
	ano[4] = "Abril"
	ano[5] = "Maio"
	ano[6] = "Junho"
	ano[7] = "Julho"
	ano[8] = "Agosto "
	ano[9] = "Setembro"
	ano[10] = "Outubro"
	ano[11] = "Novembro"
	ano[12] = "Dezembro"

	fmt.Println(ano[1], ano[12])
	fmt.Println("Tamanho: ", len(ano))
}
