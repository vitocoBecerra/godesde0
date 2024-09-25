package variables

import "fmt"

func MostrarNumeros() {
	var intcomun int
	intd32 := int32(10)
	glosa := "este es un string"

	fmt.Println("int comun = ", intcomun)
	fmt.Println("int 32 ", intd32)
	fmt.Println("Glosa es = ", glosa)
	fmt.Printf("hola [%s]\n", glosa)

}
