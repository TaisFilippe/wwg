package main

import (
	"fmt"
)

func main() {
	//Crie uma slice de tamanho 12 usando a função make e atribua os meses a cada uma dos elementos individualmente.
	//Printe essa slice mostrando todos os seus elementos.
	//Printe tambem uma fatia dessa slice do índice 2 ao 8.

	mes := make([]string, 12)

	mes[0] = "01"
	mes[1] = "02"
	mes[2] = "03"
	mes[3] = "04"
	mes[4] = "05"
	mes[5] = "06"
	mes[6] = "07"
	mes[7] = "08"
	mes[8] = "09"
	mes[9] = "10"
	mes[10] = "11"
	mes[11] = "12"

	fmt.Println(mes)

	fmt.Println(mes[1:4])
}
