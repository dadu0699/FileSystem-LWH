package ejecucion

import (
	"Sistema-de-archivos-LWH/analisis/token"
	"Sistema-de-archivos-LWH/disco/acciones"
	"Sistema-de-archivos-LWH/disco/sarchivos"
	"Sistema-de-archivos-LWH/util"
	"Sistema-de-archivos-LWH/util/grafica"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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
		parser("MOUNT")
		mount()
	case "UNMOUNT":
		parser("UNMOUNT")
		listadoIDN()
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
		parser("REP")
		rep()
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

	if _, err := os.Stat(ruta); err == nil {
		fmt.Println(">> ¿Esta seguro de que desea eliminar el disco de forma permanente? (S)")
		fmt.Print(">> ")
		if str := util.LecturaTeclado(); strings.EqualFold(str, "S") {
			acciones.EliminarDisco(ruta)
			panic(">> Disco eliminado")
		}
	} else {
		msg := ">> " + fmt.Sprintf("%s", err)
		panic(msg)
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
			if preAnalisis.GetTipo() == "SIMBOLO_MENOS" {
				parser("SIMBOLO_MENOS")
			}

			i, _ := strconv.ParseInt(preAnalisis.GetValor(), 10, 64)
			addT = i

			if listaTokens[index-1].GetTipo() == "SIMBOLO_MENOS" {
				addT *= -1
			}
			parser("ENTERO")
		}
	}

	if addT == 0 && delelteS == "" {
		acciones.CrearParticion(tamanio, ruta, nombre, unidad, tipo, fit)
	} else if addT != 0 {
		acciones.CambiarTamanio(addT, ruta, nombre, unidad)
	} else if delelteS != "" {
		fmt.Println(">> ¿Esta seguro de que desea eliminar la partición? (S)")
		fmt.Print(">> ")
		if str := util.LecturaTeclado(); strings.EqualFold(str, "S") {
			acciones.EliminarParticion(ruta, nombre, delelteS)
		}
	}
}

func mount() {
	path := ""
	name := ""
	parametros := 0

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "PATH":
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			path = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("CADENA O RUTA")

		case "NAME":
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			name = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("ID")

			parametros++
		}
	}

	if parametros > 0 {
		sarchivos.Montar(path, name)
	} else {
		sarchivos.MostrarMount()
	}
}

func listadoIDN() {
	switch preAnalisis.GetTipo() {
	case "SIMBOLO_MENOS":
		parser("SIMBOLO_MENOS")
		idn := preAnalisis.GetValor()
		if !strings.EqualFold(string(idn[0]), "I") &&
			!strings.EqualFold(string(idn[1]), "D") &&
			!unicode.IsDigit(rune(idn[2])) {
			panic(">> SE ESPERABA EL PARAMETRO IDn")
		}
		parser("ID")
		parser("SIMBOLO_MENOS")
		parser("SIMBOLO_MAYOR")
		id := strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
		sarchivos.Desmontar(id)
		parser("ID")
		listadoIDN()
	}
}

func rep() {
	id := ""
	path := ""
	nombre := ""
	// ruta := ""

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			id = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("ID")

		case "NAME":
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			nombre = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("ID")

		case "PATH":
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			path = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("CADENA O RUTA")

		case "RUTA":
			parser("RUTA")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			// ruta = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("CADENA O RUTA")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -NAME | -PATH | -RUTA)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	posFinalPath := strings.LastIndex(path, "/")
	carpetas := ""
	for i := 0; i <= posFinalPath; i++ {
		carpetas += string(path[i])
	}
	carpetas = strings.ReplaceAll(carpetas, " ", "_")

	posEXT := strings.LastIndex(path, ".")
	fileName := ""
	for i := posFinalPath + 1; i < posEXT; i++ {
		fileName += string(path[i])
	}

	ext := ""
	for i := posEXT + 1; i < len(path); i++ {
		ext += string(path[i])
	}
	filePath := sarchivos.ObtenerPath(id)

	switch strings.ToUpper(nombre) {
	case "MBR":
		grafica.TablaDisco(filePath, carpetas, fileName, ext)
	case "DISK":
		grafica.TablaDisco(filePath, carpetas, fileName, ext)
	}
}

func parser(tipo string) {
	if preAnalisis.GetTipo() != "EOF" {
		index++
		preAnalisis = listaTokens[index]
	}
}
