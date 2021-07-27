package main

import (
	"fmt"
)

func main() {

	// Dada uma slice de string que representa uma lista de compras, use o for range para percorre-la
	// Printe cada um dos itens da lista ao lado da sua ordem de inserção.

	listaDeCompra := []string{"abacate", "sabonete", "azeite", "tomate", "banana"}

	for item, valor := range listaDeCompra {

		fmt.Println(item+1, valor)

	}
}
