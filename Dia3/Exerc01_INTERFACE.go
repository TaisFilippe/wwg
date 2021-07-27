package main

import (
	"fmt"
	"math"
)

// 1) Crie tipos para representar quadrados e círculos
// 2 )Crie uma interface que descreve o comportamento de calcular a área de uma forma geométrica com a seguinte assinatura: calculeArea() float64
// 3) Implemente esse comportamento para os dois tipos criados
// 4) Depois, crie uma função que tem como parâmetro a interface que você criou e que imprime o relatório do cálculo da área da forma geométrica
// 5) Demonstre que seus tipos implementam a interface que você criou passando valores desses tipos como argumentos na chamada dessa função

// Passo 1: criar os tipos
type Quadrado struct {
	lado float64
}

// Passo 3: Atribuir comportamento
func (q Quadrado) CalculoArea() float64 {
	return math.Pow(q.lado, 2)
}

// Paaso 1
type Circulo struct {
	raio float64
}

// Passo 3
func (c Circulo) CalculoArea() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

// Passo 2: Criar interface
type Forma interface {
	calculeArea() float64
}

// Passo 4

func reportCalculo(forma Forma) {
	area := forma.calculeArea()
	fmt.Printf("O calculo da área dessa forma é: %f", area)
}

// Passo 5
func main() {

	quadrado1 := Quadrado{
		lado: 13,
	}
	circulo1 := Circulo{
		raio: 27,
	}
	reportCalculo(quadrado1)
	reportCalculo(circulo1)
}
