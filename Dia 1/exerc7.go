package main

import (
	"fmt"
)

func main() {
	//Escreva um programa que contenha um array de strings.
	//O valor de cada elemento deve ser o numero do indice escrito por extenso.
	//Print na tela o tipo do seu array e os valores de seus elementos.

	var numeros [6]string
	numeros[0] = "zero"
	numeros[1] = "um"
	numeros[2] = "dois"
	numeros[3] = "tres"
	numeros[4] = "quatro"
	numeros[5] = "cinco"

	fmt.Printf("o tipo é %T\n", numeros)
	fmt.Println(numeros)
}
