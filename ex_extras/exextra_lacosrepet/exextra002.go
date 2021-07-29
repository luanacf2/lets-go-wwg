package main

import "fmt"

type Ap struct {
	num int
	nomeProprietario string
	vagaGaragem bool
}
 func main(){
 	ap1 := Ap{
 		num: 3001,
 		nomeProprietario: "Pedro",
 		vagaGaragem: false,
	}
	 ap2 := Ap{
		 num: 3002,
		 nomeProprietario: "Julia",
		 vagaGaragem: true,
	 }
	 ap3 := Ap{
		 num: 3003,
		 nomeProprietario: "Elisa",
		 vagaGaragem: true,
	 }

	 apartamentos := []Ap{ap1,ap2,ap3}

	 for  i := range apartamentos{
	 	if apartamentos[i].vagaGaragem == true{
	 		fmt.Printf("Apartamento: \n\tNúmero: %d\n\tProprietario(a): %s\n\tPossui vaga de garagem\n", apartamentos[i].num, apartamentos[i].nomeProprietario)
		}else{
			fmt.Printf("Apartamento: \n\tNúmero: %d\n\tProprietario(a): %s\n\tNão possui vaga de garagem\n", apartamentos[i].num, apartamentos[i].nomeProprietario)
		}
	 }
 }