package main

import (
	"fmt"
)

func main() {
	// Declarar uma variavel chamada hora e atribuir um valor int a ela.
	// Usando Switch elencar cases e printar se a hora corresponde ao período da manhã, tarde e noite.

	hora := 26

	switch {

	case hora >= 6 && hora < 12:
		fmt.Printf("perído da manhã\n")
	case hora >= 12 && hora < 18:
		fmt.Printf("perído da tarde\n")
	case hora >= 18 && hora < 24:
		fmt.Printf("perído da noite\n")
	case hora >= 0 && hora < 6:
		fmt.Printf("perído da madrugada\n")
	default:
		fmt.Printf("inválido\n")
	}
}
