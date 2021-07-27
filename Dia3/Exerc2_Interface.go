package main

import "fmt"

// Crie tipos que representam diferentes animais, com atributos que façam sentido para cada um deles
// Crie uma interface que descreve o comportamento de apresentar um animal com a seguinte assinatura: Apresenta()
// Cada animal saberá como se apresentar. Sendo assim, faça com que cada um dos tipos que você criou implemente o método Apresenta(),
// que deve printar uma frase apresentando o animal e seus atributos
// Demonstre que todos os tipos implementam a interface que você criou declarando uma slice de animais e percorrendo-a com um for range que, em todas as voltas,
//chama o método Apresenta().

type Apresentador interface {
	Apresenta()
}

type Cachorro struct {
	nome  string
	porte string
	raça  string
}

func (c Cachorro) Apresenta() {
	fmt.Printf("Este é o cachorro %s ele é de porte %s e da raça %s\n", c.nome, c.porte, c.raça)
}

type Gato struct {
	nome  string
	cor   string
	idade int
}

func (g Gato) Apresenta() {
	fmt.Printf("Esse é o gato %s sua cor é %s e ele tem %d anos\n", g.nome, g.cor, g.idade)
}

func main() {

	cachorro1 := Cachorro{
		nome:  "Bartô",
		porte: "pequeno",
		raça:  "York",
	}

	gato1 := Gato{
		nome:  "Jhonny",
		cor:   "branca",
		idade: 15,
	}
	apresentadores := []Apresentador{cachorro1, gato1}

	for _, apresentador := range apresentadores {
		apresentador.Apresenta()
	}
}
