package ejercicio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func tablaMult(numero int) {

	fmt.Println("Mostrando la tabla: ", numero)

	for i := 1; i < 11; i++ {

		fmt.Printf("Resultado %d x %d es [%d]\n", numero, i, numero*i)
	}
}

func SolicitaNumero() {

	var Numero int
	var err error

	for i := 0; ; i++ {

		fmt.Print("Ingrese Numero :")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {

			Numero, err = strconv.Atoi(scanner.Text())
			if err != nil {

				fmt.Println("Error en ingreso. Reintentar...")
			} else {
				fmt.Println("Numero Ingresado : ", Numero)
				break
			}
		}

	}

	tablaMult(Numero)

}
