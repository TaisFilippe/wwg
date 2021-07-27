package main

import (
	"fmt"
	"time"
)

// Declare uma variável e atribua a ela o ano de nascimento de uma pessoa.
// Declare uma variável e atribua a ela o ano atual.
// Escreva um programa que verifique se no ano atual essa pessoa já está apta a votar e que printe na tela uma mensagem informando caso positivo e caso negativo.

func main() {
	fmt.Println("Digite o ano que vc nasceu:")

	var birthYear int
	fmt.Scan(&birthYear)

	todayYear := time.Now().Year()

	age := todayYear - birthYear

	if age >= 16 {
		fmt.Println("Pode votar")

	} else if age < 0 {
		fmt.Println("Idade inválida")

	} else //return
	{
		fmt.Println("Não pode votar")

	}
}
