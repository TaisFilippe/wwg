package main

import "fmt"

func main() {

	//Declare 5 variáveis e atribua operações relacionais a elas
	//Usando uma linha por variável, demonstre o valor e o tipo de cada uma delas.

	a := 15 >= 15
	b := 100 < 1000
	c := 10 == 18
	d := 5 != 5
	e := 89 > 0 && 50 > 0

	fmt.Printf("o tipo de a é %T e seu valor é %v\n", a, a)
	fmt.Printf("o tipo de b é %T e seu valor é %v\n", b, b)
	fmt.Printf("o tipo de c é %T e seu valor é %v\n", c, c)
	fmt.Printf("o tipo de d é %T e seu valor é %v\n", d, d)
	fmt.Printf("o tipo de e é %T e seu valor é %v\n", e, e)
}
