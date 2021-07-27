package main

import (
	"fmt"
)

var nome = "Maria"
var hora = 15

func main() {
	cumprimento := GereCumprimento(nome, hora)
	fmt.Println(cumprimento)
}

func GereCumprimento(nome string, hora int) string {

	cumprimento := ""
	switch {

	case hora >= 6 && hora < 12:
		cumprimento = "Bom dia"
	case hora >= 12 && hora < 18:
		cumprimento = "Bom tarde"
	case hora >= 18 && hora < 24:
		cumprimento = "Bom noite"
	case hora >= 0 && hora < 6:
		cumprimento = "Boa madrugada"
	default:
		cumprimento = "Ola"
	}

	frase := fmt.Sprintf("%s, %s!", cumprimento, nome)
	return frase

}
