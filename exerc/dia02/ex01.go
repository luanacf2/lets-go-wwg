package main

import "fmt"

type Pessoa struct {
	Nome string
	idade int
	corPreferida string
}
func main() {
	maria := Pessoa{"Maria", 25, "roxo"}
	pedro := Pessoa{"Pedro", 45, "verde"}
	laisa := Pessoa{"Laisa", 37, "vermelho"}

	fmt.Printf("%s: tem %d anos e sua cor preferida é %s. \n", maria.Nome, maria.idade, maria.corPreferida)
	fmt.Printf("%s: tem %d anos e sua cor preferida é %s. \n", pedro.Nome, pedro.idade, pedro.corPreferida)
	fmt.Printf("%s: tem %d anos e sua cor preferida é %s. \n", laisa.Nome, laisa.idade, laisa.corPreferida)

}