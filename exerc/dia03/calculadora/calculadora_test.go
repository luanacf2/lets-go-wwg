package calculadora

import "testing"

func TestMultiplique(t *testing.T){
	testes := []struct {
		nome  string
		x int
		y int
		esperado  int
	}{
		{
			nome:      " 7 multiplicado por e 5 é igual a 35",
			x: 7,
			y: 5,
			esperado:  35,
		},
		{
			nome:      "3 multiplicado por 20 é igual a 60",
			x: 3,
			y: 20,
			esperado:  60,
		},
		{
			nome: "27 multiplicado por 0 é igual a 0",
			x: 27,
			y: 0,
			esperado: 0,
		},
		{
			nome: "13 multiplicado por 1000 é igual a 13000",
			x: 13,
			y: 1000,
			esperado: 13000,
		},
		{
			nome: "7 multiplicado por 3 é igual a 21",
			x: 7,
			y: 2,
			esperado: 21,
		},
	}
	for _, teste := range testes{
		t.Run(teste.nome, func(t *testing.T) {
			obtive := Multiplique(teste.x, teste.y)

			if obtive != teste.esperado{
				t.Errorf("Espero '%d', mas obtive '%d'", teste.esperado, obtive)
			}
		})
	}
	//t.Run("7 multiplicado por 5 é igual a 35", func(t *testing.T) {
	//	obtive := Multiplique(7, 5)
	//	espero := 35

	//	if obtive != espero{
	//		t.Errorf("Espero '%d', mas obtive '%d'", espero, obtive)
	//	}
	//})
	//t.Run("3 multiplicado por 20 é igual a 60", func(t *testing.T) {
	//	obtive := Multiplique(3, 20)
		//espero := 60

	//	if obtive != espero{
	//		t.Errorf("Espero '%d', mas obtive '%d'", espero, obtive)
	//	}
	//})

}
