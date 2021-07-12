package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Printf("Digite o ano em que nasceu: ")
	var nasc int
	fmt.Scanln(&nasc)
	t := time.Now()
	year:= t.Year()
	idade := year - nasc

	fmt.Println("Você tem", idade, "anos.")
}
