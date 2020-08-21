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
	} else {
		panic(">> INSTRUCCION NO ENCONTRADA\n")
	}
}

func instruccion() {
	switch preAnalisis.GetTipo() {
	case "COMENTARIO":
		parser("COMENTARIO")
	case "PAUSE":
		parser("PAUSE") // util.LecturaTeclado()
	case "MKDISK":
		parser("MKDISK")
		paramsMKDISK(false, false, false, false)
	case "RMDISK":
		parser("RMDISK")
		path()
	case "FDISK":
		parser("RMDISK")
	case "MOUNT":
	case "UNMOUNT":
	}
}

func paramsMKDISK(s bool, p bool, n bool, u bool) {
	switch preAnalisis.GetTipo() {
	case "-SIZE":
		if s == true {
			panic("'ERROR PARAMETRO DUPLICADO'")
		}
		s = true
		parser("-SIZE")
		parser("ASIGNACION")
		parser("ENTERO")
	case "-PATH":
		if p == true {
			panic("'ERROR PARAMETRO DUPLICADO'")
		}
		p = true
		path()
	case "-NAME":
		if n == true {
			panic("'ERROR PARAMETRO DUPLICADO'")
		}
		n = true
		parser("-NAME")
		parser("ASIGNACION")
		parser("ID")
		parser(".DSK")
	case "-UNIT":
		if u == true {
			panic("'ERROR PARAMETRO DUPLICADO'")
		}
		u = true
		parser("-UNIT")
		parser("ASIGNACION")
		parser("ID")
	default:
		mensajePanic("(-SIZE | -PATH | -NAME | -UNIT)")
	}

	if preAnalisis.GetTipo() != "EOF" {
		paramsMKDISK(s, p, n, u)
	} else {
		if !s || !p || !n {
			mensajePanic("PARAMETROS OBLIGATORIOS")
		}
	}
}

func path() {
	parser("-PATH")
	parser("ASIGNACION")
	switch preAnalisis.GetTipo() {
	case "RUTA":
		parser("RUTA")
	case "CADENA":
		parser("CADENA")
	default:
		mensajePanic("(RUTA | \"RUTA\")")
	}
}

func parser(tipo string) {
	if preAnalisis.GetTipo() != "EOF" {
		if preAnalisis.GetTipo() != tipo {
			mensajePanic(tipo)
		}
		index++
		preAnalisis = listaTokens[index]
	}
}

func mensajePanic(mensaje string) {
	tipo := preAnalisis.GetTipo()
	if tipo == "EOF" {
		tipo = "FIN DE INSTRUCCION"
	}
	err := ">> 'ERROR: " + tipo + " SE ESPERABA " + mensaje + "'\n"
	panic(err)
}
