package main

import (
	"fmt"
)

func main() {
	listaDeCompras := []string{"cebola", "leite", "feijão"}
	fmt.Println(listaDeCompras)

	listaDeCompras = append(listaDeCompras, "banana", "chocolate")
	fmt.Println(listaDeCompras)

	listaMae := []string{"detergente", "arrroz", "sabonete"}
	listaDeCompras = append(listaDeCompras, listaMae...)
	fmt.Println(listaDeCompras)

}
