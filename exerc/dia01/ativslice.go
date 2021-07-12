package main

import "fmt"

func main()  {
	var number = []int64{2, 8, 12, 7, 26, 13}
	var number2 = []int64{98, 31, 1, 33, 24, 3}
	number3 := append(number, number2...)
	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)


}
