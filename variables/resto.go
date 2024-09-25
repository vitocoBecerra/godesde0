package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Sueldo float32
var Boolean bool
var Fecha time.Time

func RestoVariables() {

	Nombre = "pedro Araya"
	Sueldo = 23455.75
	Boolean = true
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Sueldo)
	fmt.Println(Boolean)
	fmt.Println(Fecha)

}

func ConvertToText(numero int) (bool, string) {

	texto := strconv.Itoa(numero)
	return true, texto
}
