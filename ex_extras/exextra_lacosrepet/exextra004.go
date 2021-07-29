package main

import (
	"fmt"
	"os"
)

func main() {
	var numero, n int
	for n == 0{
		fmt.Println("Insira um número: ")
		fmt.Fscanf(os.Stdin, "%d", &numero)
		if numero % 2 == 0{
			n++
		}
	}
	fmt.Println("Agora sim podemos dividir igualmente entre mim e você!")
}
