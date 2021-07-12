package main

import "fmt"

func main() {
	cores := map[string]string{
		"cinza": "#424242",
		"roxo":  "#a378f8",
	}

	 fmt.Println(cores)

	delete(cores, "cinza")

	fmt.Println(cores)

}