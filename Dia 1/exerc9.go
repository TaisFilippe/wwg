package main

import (
	"fmt"
)

func main() {
	//Declare duas slices de int e preencha 6 posições de cada uma.
	//Some as slices, formando uma terceira slice.
	//printe as tres.

	numerosX := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numerosX)

	numerosY := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(numerosY)

	numerosZ := append(numerosX, numerosY...)
	fmt.Println(numerosZ)

	parte := numerosX[0:3]
	fmt.Println(parte)
}
