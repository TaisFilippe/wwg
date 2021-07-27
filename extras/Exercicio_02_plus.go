package main

import (
	"fmt"
	"time"
)

// 1) Dado o ano em que a pessoa nasceu, calcule quantos anos ela tem no ano atual (para fins de arredondamento, considere que ela já fez aniversário no ano atual).
// 2) Como podemos pegar a informação do ano diretamente do console.

func main() {
	//fmt.Println("Digite o ano que vc nasceu:")

	//var birthdayYear int
	// fmt.Scan(&birthdayYear)
	fmt.Println("Digite o dia que vc nasceu:")
	var birthDay int
	fmt.Scan(&birthDay)

	fmt.Println("Digite o mês que vc nasceu:")
	var birthMonth time.Month
	fmt.Scan(&birthMonth)

	fmt.Println("Digite o ano que vc nasceu:")
	var birthYear int
	fmt.Scan(&birthYear)

	birth := time.Date(birthYear, birthMonth, birthDay, 0, 0, 0, 0, time.UTC)
	fmt.Println(birth)

	today := time.Now()
	fmt.Println(today)

	//age := todayYear - birthdayYear

	//fmt.Println(age)

}
