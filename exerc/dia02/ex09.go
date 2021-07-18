package main

import "fmt"

func main(){
	hora:= 0
	for hora < 24 {
		for minuto := 0; minuto < 60; minuto ++{
			fmt.Printf("%.2d:%.2d\n", hora, minuto)
		}
		hora++
	}
}
