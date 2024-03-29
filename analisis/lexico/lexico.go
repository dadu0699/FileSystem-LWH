package lexico

import (
	"FileSystem-LWH/analisis/errort"
	"FileSystem-LWH/analisis/token"
	"strings"
	"unicode"
)

// Variables Globales
var auxiliar strings.Builder
var estado int
var idToken int
var idError int
var fila int
var columna int
var listaErrores []errort.ErrorT // var listaErrores = make([]errort.ErrorT, 0)
var listaTokens []token.Token

func inicializar() {
	auxiliar.Reset()
	estado = 0
	idToken = 0
	idError = 0
	fila = 1
	columna = 1
	listaErrores = nil
	listaTokens = nil
}

// Scanner realiza el analisis lexico
func Scanner(entrada string) ([]token.Token, []errort.ErrorT) {
	inicializar()
	caracter := ""
	entrada += "\n#"

	for i := 0; i < len(entrada); i++ {
		// caracter = string([]rune(entrada)[i])
		caracter = string(entrada[i])
		// fmt.Println(caracter)

		switch estado {
		case 0:
			if esLetra(caracter) { // Palabras Reservadas o identificadores
				estado = 1
				auxiliar.WriteString(caracter)
			} else if esDigito(caracter) { // Digitos enteros o flotantes
				estado = 2
				auxiliar.WriteString(caracter)
			} else if caracter == "\"" { // Cadenas
				estado = 4
				auxiliar.WriteString(caracter)
			} else if caracter == "\\" { // Siguiente Linea
				estado = 5
				auxiliar.WriteString(caracter)
			} else if caracter == "/" { // Rutas
				estado = 8
				auxiliar.WriteString(caracter)
			} else if esEspacio(caracter) { // Espacios
				estado = 0
				auxiliar.Reset()
				// Cambio de fila y reinicio de columnas en los saltos de linea
				if caracter == "\n" {
					columna = 1
					fila++
				}
			} else if caracter == "#" { // Comentarios
				if i == len(entrada)-1 {
					// fmt.Println("Análisis léxico completado")
				} else {
					estado = 9
					auxiliar.WriteString(caracter)
				}
			} else if !agregarSimbolo(caracter) {
				agregarError(caracter)
				estado = 0
			}
		case 1:
			if esLetra(caracter) || esDigito(caracter) || caracter == "_" {
				// estado = 1
				auxiliar.WriteString(caracter)
			} else {
				agregarComando()
				i--
			}
		case 2:
			if esDigito(caracter) {
				// estado = 2
				auxiliar.WriteString(caracter)
			} else if caracter == "." {
				estado = 6
				auxiliar.WriteString(caracter)
			} else {
				agregarToken("ENTERO")
				i--
			}
		case 4:
			if caracter != "\"" {
				// estado = 4
				auxiliar.WriteString(caracter)
			} else {
				auxiliar.WriteString(caracter)
				agregarToken("CADENA")
			}
		case 5:
			if caracter == "*" {
				auxiliar.WriteString(caracter)
				agregarToken("CONTINUAR")
			} else {
				agregarError(string(entrada[i-1]))
				auxiliar.Reset()
				estado = 0
				i--
			}
		case 6:
			if esDigito(caracter) {
				// estado = 6
				auxiliar.WriteString(caracter)
			} else {
				if esDigito(string(entrada[i-1])) {
					agregarToken("DECIMAL")
				} else {
					agregarError(string(entrada[i-1]))
					auxiliar.Reset()
					estado = 0
				}
				i--
			}
		case 8:
			if !esEspacio(caracter) {
				// estado = 8
				auxiliar.WriteString(caracter)
			} else {
				agregarToken("RUTA")
			}
		case 9:
			if caracter != "\n" {
				estado = 9
				auxiliar.WriteString(caracter)
			} else {
				auxiliar.WriteString(caracter)
				agregarToken("COMENTARIO")
				columna = 1
				fila++
			}
		}
		columna++
	}
	return listaTokens, listaErrores
}

func agregarSimbolo(caracter string) bool {
	switch caracter {
	case "?":
		auxiliar.WriteString(caracter)
		agregarToken("SIMBOLO_INTERROGACION")
		return true
	case "*":
		auxiliar.WriteString(caracter)
		agregarToken("SIMBOLO_ASTERISCO")
		return true
	case "-":
		auxiliar.WriteString(caracter)
		agregarToken("SIMBOLO_MENOS")
		return true
	case ">":
		auxiliar.WriteString(caracter)
		agregarToken("SIMBOLO_MAYOR")
		return true
	case ".":
		auxiliar.WriteString(caracter)
		agregarToken("SIMBOLO_PUNTO")
		return true
	default:
		return false
	}
}

func agregarComando() {
	switch strings.ToLower(auxiliar.String()) {
	case "cat":
		agregarToken("CAT")
	case "chgrp":
		agregarToken("CHGRP")
	case "chmod":
		agregarToken("CHMOD")
	case "chown":
		agregarToken("CHOWN")
	case "cp":
		agregarToken("CP")
	case "exec":
		agregarToken("EXEC")
	case "edit":
		agregarToken("EDIT")
	case "fdisk":
		agregarToken("FDISK")
	case "find":
		agregarToken("FIND")
	case "login":
		agregarToken("LOGIN")
	case "logout":
		agregarToken("LOGOUT")
	case "mkdir":
		agregarToken("MKDIR")
	case "mkdisk":
		agregarToken("MKDISK")
	case "mkfile":
		agregarToken("MKFILE")
	case "mkfs":
		agregarToken("MKFS")
	case "mkgrp":
		agregarToken("MKGRP")
	case "mkusr":
		agregarToken("MKUSR")
	case "mount":
		agregarToken("MOUNT")
	case "mv":
		agregarToken("MV")
	case "numero": // Revisar
		agregarToken("NUMERO")
	case "pause":
		agregarToken("PAUSE")
	case "ren":
		agregarToken("REN")
	case "rep":
		agregarToken("REP")
	case "rm":
		agregarToken("RM")
	case "rmdisk":
		agregarToken("RMDISK")
	case "rmgrp":
		agregarToken("RMGRP")
	case "rmusr":
		agregarToken("RMUSR")
	case "unmount":
		agregarToken("UNMOUNT")

	// PARAMETROS
	case "add":
		agregarToken("ADD")
	case "cont":
		agregarToken("CONT")
	case "delete":
		agregarToken("DELETE")
	case "dest":
		agregarToken("DEST")
	case "file":
		agregarToken("FILEN")
	case "fit":
		agregarToken("FIT")
	case "grp":
		agregarToken("GRP")
	case "id":
		agregarToken("IDN")
	case "name":
		agregarToken("NAME")
	case "nombre": // Revisar
		agregarToken("NOMBRE")
	// case "p":
	// agregarToken("P")
	case "path":
		agregarToken("PATH")
	case "pwd":
		agregarToken("PWD")
	case "r":
		agregarToken("R")
	case "rf":
		agregarToken("RF")
	case "ruta":
		agregarToken("RUTA")
	case "size":
		agregarToken("SIZE")
	case "tipo": // Revisar
		agregarToken("TIPO")
	case "type":
		agregarToken("TYPE")
	case "ugo":
		agregarToken("UGO")
	case "unit":
		agregarToken("UNIT")
	case "usr":
		agregarToken("USR")

	// EXTENCIONES
	case "mia":
		agregarToken("MIA")
	case "dsk":
		agregarToken("DSK")
	default:
		agregarToken("ID")
	}
}

func agregarToken(tipo string) {
	idToken++
	token := new(token.Token)
	token.Inicializar(idToken, fila, columna-len(auxiliar.String()), tipo, auxiliar.String())
	listaTokens = append(listaTokens, *token)
	auxiliar.Reset()
	estado = 0
}

func agregarError(caracter string) {
	idError++
	errorT := new(errort.ErrorT)
	errorT.Inicializar(idError, fila, columna, "Léxico", caracter, "Patrón desconocido")
	listaErrores = append(listaErrores, *errorT)
}

func esLetra(caracter string) bool {
	return unicode.IsLetter(rune(caracter[0]))
}

func esDigito(caracter string) bool {
	return unicode.IsDigit(rune(caracter[0]))
}

func esEspacio(caracter string) bool {
	return unicode.IsSpace(rune(caracter[0]))
}

// ListaTokens retorna el listado de tokens
func ListaTokens() []token.Token {
	return listaTokens
}

// ListaErrores retorna el listado de tokens
func ListaErrores() []errort.ErrorT {
	return listaErrores
}
