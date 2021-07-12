package main

import "fmt"

func main()  {
	signos := make([]string, 12)
	signos[0] = "Áries"
	signos[1] = "Touro"
	signos[2] = "Gêmeos"
	signos[3] = "Câncer"
	signos[4] = "Leão"
	signos[5] = "Virgem"
	signos[6] = "Libra"
	signos[7] = "Escorpião"
	signos[8] = "Sagitário "
	signos[9] = "Capricórnio"
	signos[10] = "Aquário"
	signos[11] = "Peixes"

	fmt.Println(signos)
	fmt.Println(signos[2:8])
}
