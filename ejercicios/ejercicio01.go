package ejercicios

import (
	"fmt"
	"strconv"
)

func ConvertiraEntero(cadena string) (int, string) {
	numero, err := strconv.Atoi(cadena)
	if err != nil {
		return 0, "Error al convertir"
	}
	var resultado string

	if numero > 100 {
		resultado = "Es mayor a 100"
	} else {
		resultado = "Es menor a 100"
	}

	fmt.Println(resultado)

	return numero, resultado
}
