package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func (c Circulo) CalculoArea()float64{
	return math.Pi * math.Pow(c.raio, 2)
}

func (c Circulo) CalculoPerimetro()float64{
	return 2 * math.Pi * c.raio
}
func main(){
	circulo := Circulo{
		raio: 8,
	}
	area  := circulo.CalculoArea()
	perimetro := circulo.CalculoPerimetro()
	fmt.Printf("Círculo: \n\tRaio: %.2f\n\tÁrea: %.2f \n\tPerímetro: %.2f", circulo.raio, area, perimetro)
}

