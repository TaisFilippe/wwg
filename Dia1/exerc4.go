package main

import (
	"fmt"
)

func main() {

	// Declare duas variáveis do tipo string, uma vai ser o seu nome e outra a sua cor preferida
	// Printe na tela utilizando o fmt.Printf()e o especificador de formato %s.

	nome := "Tais"

	cor := "tiffany"

	fmt.Printf("Meu nome é %s e minha cor preferiada é %s\n", nome, cor)

	cor = "rosa"

	fmt.Printf("Meu nome é %s e agora minha cor preferiada é %s\n", nome, cor)
}
