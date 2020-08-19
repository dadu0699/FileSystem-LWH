package main

import (
	"Sistema-de-archivos-LWH/analisis/lexico"
	"fmt"
)

func main() {
	//fmt.Println(new(token.Token))
	// fmt.Println(new(errort.ErrorT))
	lexico.Scanner("exec –path->/home/Desktop/calificacion.mia")
	fmt.Println(lexico.ListaTokens())
	fmt.Println(lexico.ListaErrores())
}
