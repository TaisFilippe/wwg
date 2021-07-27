package main

import (
	"fmt"
)

func main() {

	//Utilizando uma sintaxe da estrutura for diferente da usada no exercício anterior, printe na tela todas as horas do dia (usando o formato de 24 horas).

	hora := 0

	for hora < 24 {
		fmt.Printf("%.2d:00\n", hora)
		hora++
	}
}
