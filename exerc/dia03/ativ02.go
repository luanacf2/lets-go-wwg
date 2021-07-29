package main

import "fmt"

type Apresentador interface {
	Apresenta()
}

type Humano struct {
	nome string
	corPreferida string
}

func (h Humano) Apresenta(){
	fmt.Printf("Olá! Meu nome é %s e minha cor preferida é %s\n", h.nome, h.corPreferida)
}

type Cachorro struct {
	nome string
	brinquedoPreferido string
}

func (c Cachorro) Apresenta(){
	fmt.Printf("AuAu! Meu nome é %s e meu brinquedo peferido é é %s\n", c.nome, c.brinquedoPreferido)
}

type Gato struct {
	nome string
	saborSache string
}

func (g Gato) Apresenta(){
	fmt.Printf("Miau! Meu nome é %s e meu sabor de sachê preferido é %s\n", g.nome, g.saborSache)

}

type Calopsita struct {
	nome string
	tamanhoTopete string
}
func (s Calopsita) Apresenta(){
	fmt.Printf("brr! Meu nome é %s e o tamanho do meu topete é %s\n", s.nome, s.tamanhoTopete)
}

func main(){
	humano1 := Humano{
		nome: "Shakira",
		corPreferida: "laranja",
	}

	cachorro1 := Cachorro{
		nome: "Champion",
		brinquedoPreferido: "ossinho",
	}

	gato1 := Gato{
		nome: "Fifi",
		saborSache: "cordeiro",
	}

	gato2 := Gato{
		nome: "Mingau",
		saborSache: "frango",
	}

	calopsita1 := Calopsita{
		nome: "Calô",
		tamanhoTopete: "médio",
	}

	apresentadores := []Apresentador{humano1, cachorro1, gato1, gato2, calopsita1}

	for _, apresentador := range apresentadores{
		apresentador.Apresenta()
	}

	//humano1.Apresenta()
	//cachorro1.Apresenta()
	//gato1.Apresenta()
}
