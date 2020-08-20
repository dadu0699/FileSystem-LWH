package lexico

import (
	"Sistema-de-archivos-LWH/analisis/errort"
	"Sistema-de-archivos-LWH/analisis/token"
	"fmt"
	"strings"
	"unicode"
)

// Variables Globales
var auxiliar strings.Builder
var estado int = 0
var idToken int = 0
var idError int = 0
var fila = 1
var columna = 1
var listaErrores []errort.ErrorT
var listaTokens []token.Token

// Scanner realiza el analisis lexico
func Scanner(entrada string) {
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
			} else if caracter == "-" { // Parametros
				estado = 3
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
					fmt.Println("Análisis léxico completado")
				} else {
					estado = 9
					auxiliar.WriteString(caracter)
				}
			} else if caracter == "." { // Extension
				estado = 10
				auxiliar.WriteString(caracter)
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
		case 3:
			if esLetra(caracter) {
				estado = 7
				auxiliar.WriteString(caracter)
			} else if caracter == ">" {
				auxiliar.WriteString(caracter)
				agregarToken("ASIGNACION")
			} else {
				agregarError(caracter)
				auxiliar.Reset()
				estado = 0
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
				agregarError(caracter)
				auxiliar.Reset()
				estado = 0
				i--
			}
		case 6:
			if esDigito(caracter) {
				// estado = 6
				auxiliar.WriteString(caracter)
			} else {
				agregarToken("DECIMAL")
				i--
			}
		case 7:
			if esLetra(caracter) {
				// estado = 7
				auxiliar.WriteString(caracter)
			} else {
				agregarParametro()
				i--
			}
		case 8:
			if !esEspacio(caracter) {
				// estado = 4
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
		case 10:
			if esLetra(caracter) {
				// estado = 10
				auxiliar.WriteString(caracter)
			} else {
				agregarExtension()
				i--
			}
		}
		columna++
	}
}

func agregarSimbolo(caracter string) bool {
	switch caracter {
	case "?":
		auxiliar.WriteString(caracter)
		agregarToken("SIMBOLO_INTERROGACION")
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
	default:
		agregarToken("ID")
	}
}

func agregarParametro() {
	switch strings.ToLower(auxiliar.String()) {
	case "-add":
		agregarToken("-ADD")
	case "-cont":
		agregarToken("-CONT")
	case "-delete":
		agregarToken("-DELETE")
	case "-dest":
		agregarToken("-DEST")
	case "-filen":
		agregarToken("-FILEN")
	case "-fit":
		agregarToken("-FIT")
	case "-grp":
		agregarToken("-GRP")
	case "-id":
		agregarToken("-ID")
	case "-name":
		agregarToken("-NAME")
	case "-nombre": // Revisar
		agregarToken("-NOMBRE")
	case "-p":
		agregarToken("-P")
	case "-path":
		agregarToken("-PATH")
	case "-pwd":
		agregarToken("-PWD")
	case "-r":
		agregarToken("-R")
	case "-rf":
		agregarToken("-RF")
	case "-ruta":
		agregarToken("-RUTA")
	case "-size":
		agregarToken("-SIZE")
	case "-tipo": // Revisar
		agregarToken("-TIPO")
	case "-type":
		agregarToken("-TYPE")
	case "-ugo":
		agregarToken("-UGO")
	case "-unit":
		agregarToken("-UNIT")
	case "-usr":
		agregarToken("-USR")
	default:
		agregarError(auxiliar.String())
		auxiliar.Reset()
		estado = 0
	}
}

func agregarExtension() {
	switch auxiliar.String() {
	case ".mia":
		agregarToken(".MIA")
	case ".dsk":
		agregarToken(".DSK")
	default:
		agregarError(auxiliar.String())
		auxiliar.Reset()
		estado = 0
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
