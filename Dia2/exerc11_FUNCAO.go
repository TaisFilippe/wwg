package main

import (
	"fmt"
)

//Dado um número A e um número B, escreva uma função que some ambos e retorne a soma.

func main() {

	soma := Some(4, 19)
	fmt.Println(soma)
}

func Some(a, b int) int {

	soma := a + b
	return soma

}
