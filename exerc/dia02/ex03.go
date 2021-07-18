package main

import "fmt"

func main(){
 	num := -20

 	if num > 0{
 		fmt.Println("O número é positivo!")
		return
	}else if num < 0 {
		fmt.Println("O número é negativo!")
	}else{
		fmt.Println("O número é 0!")
	}
}

