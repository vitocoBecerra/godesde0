package main

import (
	"fmt"
	"godesde0/variables"
)

func main() {

	variables.MostrarNumeros()

	variables.RestoVariables()

	boolen, texto := variables.ConvertToText(35098)

	fmt.Println(boolen)
	fmt.Println(texto)

}
