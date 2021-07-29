package main

import "fmt"

type Conjunto[]int

func (c Conjunto) Soma()int{
	var soma int
	for _, n := range c{
		soma += n
	}
	return soma
}

func (c Conjunto) Media()float64{
	soma := c.Soma()
	media := float64(soma)/ float64(len(c))
	return media
}

func main(){
	conjunto1 := Conjunto{3, 8, 16, 60, 35}
	conjunto2 := Conjunto{8, 6, 1, 25, 40}
	 var conjuntos = []Conjunto{conjunto1,conjunto2}

	 for _, valor := range  conjuntos{
	 	soma := valor.Soma()
	 	media := valor.Media()
	 	fmt.Printf("nosso conjunto tem os elementos: %v\n\tsoma: %d\n\tmédia: %f\n", valor, soma, media)

	 }
}