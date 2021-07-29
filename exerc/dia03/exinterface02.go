package main

import "fmt"

type Apresentar interface {
	Apresenta()
}

type Aves struct {
	nome string
	corDasPenas string
	idade int
}

type Macacos struct {
	nome string
	genero string
	raça string
}
 func (a Aves) Apresenta(){
 	fmt.Printf("Meu nome é %s, a cor das minhas penas é %s e minha idade é %d\n", a.nome, a.corDasPenas, a.idade)
 }

 func (m Macacos) Apresenta(){
 	fmt.Printf("Meu nome é %s, e meu gênero é %s e minha raça é %s\n", m.nome, m.genero, m.raça)
 }

 func main(){
 	ave1 := Aves{
 		nome: "Fifi",
 		corDasPenas: "amarelo",
 		idade: 1,
	}
	ave2 := Aves{
		nome: "Sansão",
		corDasPenas: "cinza",
		idade: 2,
	}
	macaco1 := Macacos{
		nome: "Mel",
		genero: "feminino",
		raça: "chimpanzé",
	}
	 macaco2 := Macacos{
		 nome: "Mico",
		 genero: "masculino",
		 raça: "mico-leão",
	 }

	 animais := []Apresentar{ave1, ave2, macaco1, macaco2}

	 for _, anim := range animais{
	 	anim.Apresenta()
	 }
 }