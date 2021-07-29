package main

import "fmt"

func main() {
	for contador := 0; contador < 10; contador++{
		for linha := 0; linha <= contador; linha++{
			fmt.Printf("%d ", linha+1)
		}
		fmt.Printf("\n")
	}
}