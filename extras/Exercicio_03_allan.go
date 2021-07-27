package main

import "fmt"

// Escreva um programa que lista o nome dos países e quantas vezes eles aparecem no nosso map.

// Passos:
// 1)Criar um mapa com chave do tipo string e valor do tipo string (map[string]string) onde a chave seja o nome da cidade e o valor o nome do país.
// 2)Percorrer o mapa do item 1 acumulando em outro mapa a frequência de aparições do país. Esse segundo mapa terá tipo map[string]int, sendo a chave o nome do país e o valor a quantidade de menções.
// 3)Printe na tela os valores do último mapa.

func main() {

	meuMapa := make(map[string]int)

	meuMapa["Paranagua"] = 1
	meuMapa["Lyon"] = 3
	meuMapa["Curitiba"] = 2
	meuMapa["Los Angeles"] = 7
	meuMapa["Florenca"] = 4
	meuMapa["São Paulo"] = 1
	meuMapa["Veneza"] = 6

	outroMapa := make(map[int]string)

	for cidade, valor := range meuMapa {
		antigaCidade, existe := outroMapa[valor]
		if !existe {
			outroMapa[valor] = cidade
		} else {
			outroMapa[valor] = antigaCidade + " - " + cidade
		}
	}

	fmt.Printf("%#v\n", outroMapa)

}
