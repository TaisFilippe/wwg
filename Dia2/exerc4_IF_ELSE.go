package main

import (
	"fmt"
)

// Faça um programa que, dada a idade de uma pessoa, verifique se ela é menor de idade, maior de idade ou idosa.
// Leve em consideração os seguintes valores:
// Menor de idade = abaixo de 18;
//Maior de idade = entre 18 e 60;
//Idoso = acima de 60.

func main() {

	idade := 47

	if idade >= 60 {
		fmt.Printf("a pessoa tem idade %d e é idosa\n", idade)

	} else if idade >= 18 && idade < 60 {
		fmt.Printf("a pessoa tem idade %d e é maior de idade\n", idade)

	} else if idade < 0 {
		fmt.Println("o valor informado é invalido\n", idade)

	} else {
		fmt.Printf("a pessoa tem idade %d e é menor de idade\n", idade)
	}

}
