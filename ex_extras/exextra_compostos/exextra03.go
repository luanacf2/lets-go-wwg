package main

import "fmt"

func main(){
	paises := map[string]string{
		"Londres" : "Reino Unido",
		"Tóquio" : "Japão",
		"Nova York" : "EUA",
		"Miami" : "EUA",
		"Barcelona" : "Espanha",
		"Madrid" : "Espanha",
		"Veneza" : "Itália",
		"Berlim" : "Alemanha",
		"Milão" : "Itália",
		"Orlando" : "EUA",
		"Munique" : "Alemanha",
		"Rio de Janeiro" : "Brasil",
		"Buenos Aires" : "Argentina",
		"Toronto" : "Canadá",
		"Cancun" : "México",
	}
	frequencia := make(map[string]int)

	for _, valor := range paises{
		frequencia[valor] += 1
	}

	fmt.Println(frequencia)

}
