/*Crear una funcion publica que pida un numero por teclado
valide si hay un error o no y que si hay error que vuelva a pedirlo
en base al numero recibido crear la tabla numerica del mismo
de 1 a 10 y la muestre por pantalla
*/

package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaNum() {

	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error

	for {

		fmt.Println("Ingrese un numero")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {

		numero = numero * i
		fmt.Printf("%d x %d = %d ", numero, i, i*numero)
	}
}
