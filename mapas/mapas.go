package mapas

import "fmt"

func MostrarMapas() {

	paises := make(map[string]string)

	paises["Mexico"] = "D.F"
	paises["Chile"] = "Santiago"

	fmt.Println(paises)
	fmt.Println(paises["Chile"])
}
