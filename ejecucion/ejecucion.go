package ejecucion

import (
	"Sistema-de-archivos-LWH/analisis/token"
	"Sistema-de-archivos-LWH/disco/acciones"
	"Sistema-de-archivos-LWH/util"
	"fmt"
	"strconv"
	"strings"
)

// Variables globales
var index int
var preAnalisis token.Token
var listaTokens []token.Token

// Ejecutar cada una de las instrucciones
func Ejecutar(listadoAnalisisLexico []token.Token) {
	index = 0
	listaTokens = listadoAnalisisLexico

	token := new(token.Token)
	token.Inicializar(0, 0, 0, "EOF", "EOF")
	listaTokens = append(listaTokens, *token)
	preAnalisis = listaTokens[index]

	iniciar()
}

func iniciar() {
	switch preAnalisis.GetTipo() {
	case "PAUSE":
		parser("PAUSE")
		util.LecturaTeclado()

	case "MKDISK":
		parser("MKDISK")
		mkdisk()

	case "RMDISK":
		parser("RMDISK")
		rmdisk()

	case "FDISK":
		parser("FDISK")
		fdisk()

	case "MOUNT":

	case "UNMOUNT":
	}
}

func mkdisk() {
	var tamanio int64
	ruta := ""
	nombre := ""
	unidad := ""

	for preAnalisis.GetTipo() != "EOF" {
		switch preAnalisis.GetTipo() {
		case "-SIZE": // OBLIGATORIO
			parser("-SIZE")
			parser("ASIGNACION")

			i, _ := strconv.ParseInt(preAnalisis.GetValor(), 10, 64)
			tamanio = i
			if tamanio <= 0 {
				panic("SOLO SE ACEPTAN NUMEROS ENTEROS POSITIVOS DE TAMAÑO")
			}
			parser("ENTERO")

		case "-PATH": // OBLIGATORIO
			parser("-PATH")
			parser("ASIGNACION")
			ruta = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("CADENA O RUTA")

		case "-NAME": // OBLIGATORIO
			parser("-NAME")
			parser("ASIGNACION")
			nombre = preAnalisis.GetValor() + ".dsk"
			parser("ID")
			parser(".DSK")

		case "-UNIT":
			parser("-UNIT")
			parser("ASIGNACION")
			unidad = preAnalisis.GetValor()
			if !strings.EqualFold(unidad, "K") &&
				!strings.EqualFold(unidad, "M") {
				panic("PARAMETRO DE UNIDAD INCORRECTO")
			}
			parser("ID")
		}
	}
	acciones.CrearDisco(tamanio, ruta, nombre, unidad)
}

func rmdisk() {
	parser("-PATH")
	parser("ASIGNACION")

	ruta := strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
	parser("CADENA O RUTA")

	fmt.Println(">> ¿Esta seguro de que desea eliminar el disco de forma permanente?")
	if str := util.LecturaTeclado(); strings.EqualFold(str, "S") {
		acciones.EliminarDisco(ruta)
		fmt.Println(">> Disco eliminado")
	}
}

func fdisk() {
	var tamanio int64
	var addT int64
	ruta := ""
	nombre := ""
	unidad := ""
	tipo := ""
	fit := ""
	delelteS := ""

	for preAnalisis.GetTipo() != "EOF" {
		switch preAnalisis.GetTipo() {
		case "-SIZE": // OBLIGATORIO
			parser("-SIZE")
			parser("ASIGNACION")
			i, _ := strconv.ParseInt(preAnalisis.GetValor(), 10, 64)
			tamanio = i
			if tamanio <= 0 {
				panic("SOLO SE ACEPTAN NUMEROS ENTEROS POSITIVOS DE TAMAÑO")
			}
			parser("ENTERO")

		case "-UNIT":
			parser("-UNIT")
			parser("ASIGNACION")
			unidad = preAnalisis.GetValor()
			if !strings.EqualFold(unidad, "B") &&
				!strings.EqualFold(unidad, "K") &&
				!strings.EqualFold(unidad, "M") {
				panic("PARAMETRO DE UNIDAD INCORRECTO")
			}
			parser("ID")

		case "-PATH": // OBLIGATORIO
			parser("-PATH")
			parser("ASIGNACION")
			ruta = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("CADENA O RUTA")

		case "-TYPE":
			parser("-TYPE")
			parser("ASIGNACION")
			tipo = preAnalisis.GetValor()
			if !strings.EqualFold(tipo, "P") &&
				!strings.EqualFold(tipo, "E") &&
				!strings.EqualFold(tipo, "L") {
				panic("PARAMETRO DE TIPO INCORRECTO")
			}
			parser("ID")

		case "-FIT":
			parser("-FIT")
			parser("ASIGNACION")
			fit = preAnalisis.GetValor()
			if !strings.EqualFold(fit, "BF") &&
				!strings.EqualFold(fit, "FF") &&
				!strings.EqualFold(fit, "WF") {
				panic("PARAMETRO DE FIT INCORRECTO")
			}
			parser("ID")

		case "-DELETE":
			parser("-DELETE")
			parser("ASIGNACION")
			delelteS = preAnalisis.GetValor()
			if !strings.EqualFold(delelteS, "Full") &&
				!strings.EqualFold(delelteS, "Fast") {
				panic("PARAMETRO DE DELETE INCORRECTO")
			}
			parser("ID")

		case "-NAME": // OBLIGATORIO
			parser("-NAME")
			parser("ASIGNACION")
			nombre = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("ID")

		case "-ADD":
			parser("-ADD")
			parser("ASIGNACION")
			i, _ := strconv.ParseInt(preAnalisis.GetValor(), 10, 64)
			addT = i
			parser("ENTERO")
		}
	}
	acciones.CrearParticion(tamanio, ruta, nombre, unidad, tipo, fit, addT, delelteS)
	acciones.Graficar(ruta)
}

func parser(tipo string) {
	if preAnalisis.GetTipo() != "EOF" {
		index++
		preAnalisis = listaTokens[index]
	}
}
