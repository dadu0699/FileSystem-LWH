package sintactico

import (
	"Sistema-de-archivos-LWH/analisis/errort"
	"Sistema-de-archivos-LWH/analisis/token"
	"fmt"
)

// Variables globales
var index int
var preAnalisis token.Token
var listaTokens []token.Token
var idError int
var errorSintactico bool
var listaErrores []errort.ErrorT

// Analizar inicio del analisis sintactico
func Analizar(listadoAnalisisLexico []token.Token) {
	index = 0
	listaTokens = listadoAnalisisLexico

	token := new(token.Token)
	token.Inicializar(0, 0, 0, "EOF", "EOF")
	listaTokens = append(listaTokens, *token)
	preAnalisis = listaTokens[index]

	idError = 0
	errorSintactico = false
	listaErrores = nil

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	inicio()
}

func inicio() {
	if preAnalisis.GetTipo() == "COMENTARIO" ||
		preAnalisis.GetTipo() == "PAUSE" ||
		preAnalisis.GetTipo() == "MKDISK" ||
		preAnalisis.GetTipo() == "RMDISK" ||
		preAnalisis.GetTipo() == "FDISK" ||
		preAnalisis.GetTipo() == "MOUNT" ||
		preAnalisis.GetTipo() == "UNMOUNT" {
		instruccion()
	}
	parser("EOF")
}

func instruccion() {
	switch preAnalisis.GetTipo() {
	case "COMENTARIO":
		parser("COMENTARIO")
	case "PAUSE":
		parser("PAUSE")
		//util.LecturaTeclado()
	case "MKDISK":
	case "RMDISK":
	case "FDISK":
	case "MOUNT":
	case "UNMOUNT":
	}
}

func parser(tipo string) {
	if preAnalisis.GetTipo() != "EOF" {
		if tipo == "EOF" {
			panic(">> 'ERROR DE INSTRUCCION'\n")
		}

		if preAnalisis.GetTipo() != tipo {
			err := ">> 'ERROR: " + preAnalisis.GetTipo() + " SE ESPERABA " + tipo + "'"
			panic(err)
		}
		index++
		preAnalisis = listaTokens[index]
	}
}
