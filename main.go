package main

import (
	"fmt"

	"github.com/DanielaGrosclaude/gonuts/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1213)

	fmt.Println(estado)
	fmt.Println(texto)

	variables.RestoVariables()

}
