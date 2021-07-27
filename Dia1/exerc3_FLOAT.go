package main

import (
	"fmt"
)

func main() {

	//Calcule o valor total de uma compra que tem 3 itens representando o preço de todos eles em float64
	//Todos os itens dessa compra precisam ser comprados em mais de uma unidade

	var pesoDoTomate, quantidadeDeGarrafasDeAzeite, unidadesDeSabonete float64
	pesoDoTomate = 2.600
	quantidadeDeGarrafasDeAzeite = 2
	unidadesDeSabonete = 7

	var precoDoTomate, precoDoAzeite, precoDoSabonete float64
	precoDoTomate = 10.3
	precoDoAzeite = 19
	precoDoSabonete = 5.80

	var valorDaCompra = (pesoDoTomate * precoDoTomate) + (quantidadeDeGarrafasDeAzeite * precoDoAzeite) + (unidadesDeSabonete * precoDoSabonete)

	fmt.Println(valorDaCompra)
}
