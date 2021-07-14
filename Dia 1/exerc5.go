package main

import (
	"fmt"
)

func main() {

	estaEnsolarado := true
	estouDisposta := true
	estouSaindoDeCasa := false

	vouAPraia := estaEnsolarado && estouDisposta && estouSaindoDeCasa

	fmt.Println(vouAPraia)
	fmt.Printf("%T", vouAPraia)
}
