package ejecucion

import (
	"Sistema-de-archivos-LWH/analisis/token"
	"Sistema-de-archivos-LWH/disco/acciones"
	"Sistema-de-archivos-LWH/util"
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
		case "-SIZE":
			parser("-SIZE")
			parser("ASIGNACION")
			i, _ := strconv.ParseInt(preAnalisis.GetValor(), 10, 64)
			tamanio = i
			parser("ENTERO")
		case "-PATH":
			parser("-PATH")
			parser("ASIGNACION")
			ruta = strings.ReplaceAll(preAnalisis.GetValor(), "\"", "")
			parser("CADENA O RUTA")
		case "-NAME":
			parser("-NAME")
			parser("ASIGNACION")
			nombre = preAnalisis.GetValor() + ".dsk"
			parser("ID")
			parser(".DSK")
		case "-UNIT":
			parser("-UNIT")
			parser("ASIGNACION")
			unidad = preAnalisis.GetValor()
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
	acciones.EliminarDisco(ruta)
}

func parser(tipo string) {
	if preAnalisis.GetTipo() != "EOF" {
		index++
		preAnalisis = listaTokens[index]
	}
}
