package main

import (
	"fmt"
	"math"
)

// Construa dois métodos: um que retorna a área e outro que retorna o perímetro de uma estrutura que representa um círculo.
// Printe a área e o perímetro de um círculo.

type Circulo struct {
	raio float64
}

func (c Circulo) Area() float64 {

	return math.Pi * math.Pow(c.raio, 2)

}

func (c Circulo) Perímetro() float64 {
	return 2 * math.Pi * c.raio
}
func main() {
	circulo1 := Circulo{
		raio: 3,
	}
	area := circulo1.Area()
	perímetro := circulo1.Perímetro()
	fmt.Printf("área: %.2f\n perímetro:%.2f\n", area, perímetro)
}
