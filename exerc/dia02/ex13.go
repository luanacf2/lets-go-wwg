package main

import (
	"fmt"
)

func main() {
	soma := Soma(20,45)
	fmt.Println(soma)
}

func Soma(a int, b int) int {
	return a +b
}