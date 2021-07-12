package main

import "fmt"

func main()  {
	a := make([]string, 6)
	a[0] = "primeiro"
	a[5] = "sexto"
	fmt.Println(a)
	fmt.Println(len(a))
	capacidade := cap(a)
	fmt.Println(capacidade)

}
