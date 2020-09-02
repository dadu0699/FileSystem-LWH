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
	case "MKFS":
	case "LOGIN":
	case "MKGRP":
	case "RMGRP":
	case "MKUSR":
	case "RMUSR":
	case "CHMOD":
	case "MKFILE":
	case "CAT":
	case "RM":
	case "EDIT":
	case "REN":
	case "MKDIR":
	case "CP":
	case "MV":
	case "FIND":
	case "CHOWN":
	case "CHGRP":
	case "LOSS":
	case "RECOVERY":
	case "REP":
	}
}

func mkdisk() {
	var tamanio int64
	ruta := ""
	nombre := ""
	unidad := ""

	for preAnalisis.GetTipo() != "EOF" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "SIZE": // OBLIGATORIO
			parser("SIZE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")

			i, _ := strconv.ParseInt(preAnalisis.GetValor(), 10, 64)
			tamanio = i
			if tamanio <= 0 {
				panic(">> SOLO SE ACEPTAN NUMEROS ENTEROS POSITIVOS")
			}
			parser("ENTERO")

		case "PATH": // OBLIGATORIO
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			ruta = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("CADENA O RUTA")

		case "NAME": // OBLIGATORIO
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			nombre = preAnalisis.GetValor() + ".dsk"
			parser("ID")
			parser("SIMBOLO_PUNTO")
			parser("DSK")

		case "UNIT":
			parser("UNIT")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			unidad = preAnalisis.GetValor()
			if !strings.EqualFold(unidad, "K") &&
				!strings.EqualFold(unidad, "M") {
				panic(">> PARAMETRO DE 'UNIDAD' INCORRECTO SE ESPERABA (K | M)")
			}
			parser("ID")
		}
	}
	acciones.CrearDisco(tamanio, ruta, nombre, unidad)
}

func rmdisk() {
	parser("SIMBOLO_MENOS")
	parser("PATH")
	parser("SIMBOLO_MENOS")
	parser("SIMBOLO_MAYOR")

	ruta := strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
	parser("CADENA O RUTA")

	fmt.Println(">> ¿Esta seguro de que desea eliminar el disco de forma permanente?")
	fmt.Print(">> ")
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
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "SIZE": // OBLIGATORIO
			parser("SIZE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			i, _ := strconv.ParseInt(preAnalisis.GetValor(), 10, 64)
			tamanio = i
			if tamanio <= 0 {
				panic(">> SOLO SE ACEPTAN NUMEROS ENTEROS POSITIVOS")
			}
			parser("ENTERO")

		case "UNIT":
			parser("UNIT")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			unidad = preAnalisis.GetValor()
			if !strings.EqualFold(unidad, "B") &&
				!strings.EqualFold(unidad, "K") &&
				!strings.EqualFold(unidad, "M") {
				panic(">> PARAMETRO DE 'UNIDAD' INCORRECTO SE ESPERABA (B | K | M)")
			}
			parser("ID")

		case "PATH": // OBLIGATORIO
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			ruta = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("CADENA O RUTA")

		case "TYPE":
			parser("TYPE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			tipo = preAnalisis.GetValor()
			if !strings.EqualFold(tipo, "P") &&
				!strings.EqualFold(tipo, "E") &&
				!strings.EqualFold(tipo, "L") {
				panic(">> PARAMETRO DE 'TIPO' INCORRECTO SE ESPERABA (P | E | L)")
			}
			parser("ID")

		case "FIT":
			parser("FIT")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			fit = preAnalisis.GetValor()
			if !strings.EqualFold(fit, "BF") &&
				!strings.EqualFold(fit, "FF") &&
				!strings.EqualFold(fit, "WF") {
				panic(">> PARAMETRO DE 'FIT' INCORRECTO SE ESPERABA (BD | FF | WF)")
			}
			parser("ID")

		case "DELETE":
			parser("DELETE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			delelteS = preAnalisis.GetValor()
			if !strings.EqualFold(delelteS, "Full") &&
				!strings.EqualFold(delelteS, "Fast") {
				panic(">> PARAMETRO DE 'DELETE' INCORRECTO SE ESPERABA (FULL | FAST)")
			}
			parser("ID")

		case "NAME": // OBLIGATORIO
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			nombre = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("ID")

		case "ADD":
			parser("ADD")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			i, _ := strconv.ParseInt(preAnalisis.GetValor(), 10, 64)
			addT = i
			parser("ENTERO")
		}
	}
	if addT == 0 && delelteS == "" {
		acciones.CrearParticion(tamanio, ruta, nombre, unidad, tipo, fit)
		acciones.Graficar(ruta)
	} else if addT > 0 {

	} else if delelteS != "" {
		acciones.EliminarParticion(ruta, nombre, delelteS)
	}
}

func parser(tipo string) {
	if preAnalisis.GetTipo() != "EOF" {
		index++
		preAnalisis = listaTokens[index]
	}
}
