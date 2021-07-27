package main

import (
	"fmt"
)

func main() {

	//Utilizando a resolução anterior;
	//Demonstre todos os minutos de um dia (formato 24 horas)
	//E printe todos os segundos de 00:00:00 até 02:59:59

	hora := 0

	for hora < 3 {

		for minuto := 0; minuto < 60; minuto++ {
			for segundo := 0; segundo < 60; segundo++ {
				fmt.Printf("%.2d:%.2d:%.2d\n", hora, minuto, segundo)
			}
		}
		hora++
	}
}
