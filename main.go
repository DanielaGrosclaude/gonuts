package main

import (
	"fmt"

	"github.com/DanielaGrosclaude/gonuts/ejercicios"
	//"runtime"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1213)

	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X. " {
		fmt.Println("Esto no es Windows, es ", os)

	} else {
		fmt.Println("Esto es ", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux.")
	case "darwin":
		fmt.Println("Esto es Darwin.")
	default:
		fmt.Printf("%s \n", os)
	}*/

	cadenaNumero := "56"
	numero, resultado := ejercicios.ConvertiraEntero(cadenaNumero)

	fmt.Println("Número:", numero)
	fmt.Println("Resultado:", resultado)

}
