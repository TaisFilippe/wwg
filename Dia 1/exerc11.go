package main

import "fmt"

func main() {

	//Crie um map chamado ano onde as chaves são os meses e o valor é o numero do mes representa
	//Printe os meses janeiro e dezembro
	//Printe o tamanho do map ano

	ano := make(map[string]int)
	ano["Jan"] = 1
	ano["Fev"] = 2
	ano["Mar"] = 3
	ano["Abr"] = 4
	ano["Mai"] = 5
	ano["Jun"] = 6
	ano["Jul"] = 7
	ano["Ago"] = 8
	ano["Set"] = 9
	ano["Out"] = 10
	ano["Nov"] = 11
	ano["Dez"] = 12
	fmt.Println(ano["Jan"], ano["Dez"])

	fmt.Println("tamanho do slice: ", len(ano))
}
