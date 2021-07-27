package main

import "fmt"

// Crie um tipo Pessoa que tenha os atributos nome, idade e cor preferida.
// Crie três variáveis do tipo pessoa.
// Printe o nome de todas as três, bem como frases informando sua idade e cores preferidas.

type Pessoa struct {
	nome         string
	idade        int
	corPreferida string
}

func main() {
	pessoa1 := Pessoa{
		nome:         "Tais",
		idade:        28,
		corPreferida: "Azul Tiffany",
	}

	pessoa2 := Pessoa{
		nome:         "Allan",
		idade:        33,
		corPreferida: "Verde",
	}

	pessoa3 := Pessoa{
		nome:         "Camila",
		idade:        31,
		corPreferida: "Rosa",
	}
	fmt.Printf("dados das pessoas: %s, %s e %s\n", pessoa1.nome, pessoa2.nome, pessoa3.nome)
	fmt.Printf("A %s tem %d anos e sua cor preferida é %s\n", pessoa1.nome, pessoa1.idade, pessoa1.corPreferida)
	fmt.Printf("O %s tem %d anos e sua cor preferida é %s\n", pessoa2.nome, pessoa2.idade, pessoa2.corPreferida)
	fmt.Printf("A %s tem %d anos e sua cor preferida é %s\n", pessoa3.nome, pessoa3.idade, pessoa3.corPreferida)
}
