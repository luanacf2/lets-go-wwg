package main

import (
	"fmt"
	"math"
)

type Calculo interface {
	calculoArea()float64

}

type Quadrado struct {
	lado float64
}

func (q Quadrado) calculoArea()float64{
	return math.Pow(q.lado, 2)
}

type Circulo struct {
	raio float64
}

func (c Circulo) calculoArea()float64{
	return  math.Pi * math.Pow(c.raio, 2)
}

func RelatorioCalc( r Calculo){
	fmt.Printf("A área desta forma geométrica é %.2f\n", r.calculoArea())
}

func main(){
	quadrado1 := Quadrado{
		lado: 13,
	}
	quadrado2 := Quadrado{
		lado: 8,
	}
	circulo1 := Circulo{
		raio: 27,
	}
	circulo2 := Circulo{
		raio: 15,
	}

	RelatorioCalc(quadrado1)
	RelatorioCalc(quadrado2)
	RelatorioCalc(circulo1)
	RelatorioCalc(circulo2)

}


