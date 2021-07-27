package main

import (
	"fmt"
)

// Declara uma variável;
// Atribui o valor de um número a ela;
// Informa se o número é positivo ou negativo.

func main() {

	numero := -10

	if numero > 0 {
		fmt.Println("o numero é positivo")

	} else if numero < 0 {
		fmt.Println("o numero é negativo")

	} else {

		fmt.Println("o numero é  0")

	}
}
