package main

import "fmt"

type Gato struct {
	nome string
	pessoaResponsavel string
	saborSachePreferido string
}

func (g Gato) Apresenta(){
	fmt.Printf("Miau! Meu nome é %s, a pessoa responável por mim é %s e meu sabor de sachê preferido é %s.\n", g.nome, g.pessoaResponsavel, g.saborSachePreferido)
}
type Cachorro struct {
	nome string
	pessoaResponsavel string
	brinquedoPreferido string
}

func (c Cachorro) Apresenta(){
	fmt.Printf("AuAu! Meu nome é %s, a pessoa responável por mim é %s e meu brinquedo preferido é %s.\n", c.nome, c.pessoaResponsavel, c.brinquedoPreferido)
}

func main()  {
	gato1 := Gato{
		nome: "Mingau",
		pessoaResponsavel: "Shakira",
		saborSachePreferido: "frango",
	}
	gato2 := Gato{
		nome: "Lua",
		pessoaResponsavel: "Anitta",
		saborSachePreferido: "salmão",
	}
	cachorro1 := Cachorro{
		nome: "Belinha",
		pessoaResponsavel: "Thalia",
		brinquedoPreferido: "ossinho",
	}

	gato1.Apresenta()
	gato2.Apresenta()
	cachorro1.Apresenta()
}