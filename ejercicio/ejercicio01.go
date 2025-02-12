package ejercicio

import (
	"fmt"
	"strconv"
)

var Numero int

func ConvertirIntToString(cadena string) (int, string) {

	//var cadena string

	if Numero, _ = strconv.Atoi(cadena); Numero > 100 {
		fmt.Println("Numero mayor a 100")
		//fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("Numero menor a 100")
	}

	return Numero, cadena
}
