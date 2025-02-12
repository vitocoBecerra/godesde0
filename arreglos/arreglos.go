package arreglos

import (
	"fmt"
)

var tabla [10]int

func MuestroArreglos() {

	for i := 0; i < len(tabla); i++ {
		tabla[i] = 100 * i
	}

	for j := 0; j < len(tabla); j++ {
		fmt.Println(" el valor es ", tabla[j])
	}

}
