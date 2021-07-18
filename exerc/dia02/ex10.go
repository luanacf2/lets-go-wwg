package main

import "fmt"

func main(){
	hora:= 0
	for hora < 3 {
		minuto :=0
		for minuto < 60{
			for segundos := 0; segundos < 60; segundos++ {
				fmt.Printf("%.2d:%.2d:%.2d\n", hora, minuto,segundos)
			}
			minuto++
		}
		hora++
	}
}

