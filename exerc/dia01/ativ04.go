package main

import "fmt"

func main()  {
	var listaDeCompras = []string{"azeite", "feijão", "carne"}
	fmt.Println(listaDeCompras)
	fmt.Println(listaDeCompras[2])

	listaDeCompras = append(listaDeCompras, "leite")
	fmt.Println(listaDeCompras)

	var listaDaIrma = []string{"chocolate", "salgadinho"}
	listaDeCompras = append(listaDeCompras, listaDaIrma...)
	fmt.Println(listaDeCompras)

}
