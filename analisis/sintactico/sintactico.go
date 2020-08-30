package sintactico

import (
	"Sistema-de-archivos-LWH/analisis/token"
	"Sistema-de-archivos-LWH/ejecucion"
	"fmt"
)

// Variables globales
var index int
var preAnalisis token.Token
var listaTokens []token.Token

// Analizar inicio del analisis sintactico
func Analizar(listadoAnalisisLexico []token.Token) {
	index = 0
	listaTokens = listadoAnalisisLexico

	token := new(token.Token)
	token.Inicializar(0, 0, 0, "EOF", "EOF")
	listaTokens = append(listaTokens, *token)
	preAnalisis = listaTokens[index]

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println()
			// util.LecturaTeclado()
		}
	}()

	inicio()
	ejecucion.Ejecutar(listadoAnalisisLexico)
}

func inicio() {
	if preAnalisis.GetTipo() == "COMENTARIO" ||
		preAnalisis.GetTipo() == "PAUSE" ||
		preAnalisis.GetTipo() == "MKDISK" ||
		preAnalisis.GetTipo() == "RMDISK" ||
		preAnalisis.GetTipo() == "FDISK" ||
		preAnalisis.GetTipo() == "MOUNT" ||
		preAnalisis.GetTipo() == "UNMOUNT" ||
		preAnalisis.GetTipo() == "MKFS" ||
		preAnalisis.GetTipo() == "LOGIN" ||
		preAnalisis.GetTipo() == "LOGOUT" ||
		preAnalisis.GetTipo() == "MKGRP" ||
		preAnalisis.GetTipo() == "RMGRP" ||
		preAnalisis.GetTipo() == "MKUSR" ||
		preAnalisis.GetTipo() == "RMUSR" ||
		preAnalisis.GetTipo() == "CHMOD" ||
		preAnalisis.GetTipo() == "MKFILE" ||
		preAnalisis.GetTipo() == "CAT" ||
		preAnalisis.GetTipo() == "RM" ||
		preAnalisis.GetTipo() == "EDIT" ||
		preAnalisis.GetTipo() == "REN" ||
		preAnalisis.GetTipo() == "MKDIR" ||
		preAnalisis.GetTipo() == "CP" ||
		preAnalisis.GetTipo() == "MV" ||
		preAnalisis.GetTipo() == "FIND" ||
		preAnalisis.GetTipo() == "CHOWN" ||
		preAnalisis.GetTipo() == "CHGRP" ||
		preAnalisis.GetTipo() == "LOSS" ||
		preAnalisis.GetTipo() == "RECOVERY" ||
		preAnalisis.GetTipo() == "REP" {
		instruccion()
	} else {
		panic(">> INSTRUCCION NO ENCONTRADA")
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
		paramsMKDISK()
	case "RMDISK":
		parser("RMDISK")
		parser("SIMBOLO_MENOS")
		parser("PATH")
		parser("SIMBOLO_MENOS")
		parser("SIMBOLO_MAYOR")
		pathParams()
	case "FDISK":
		parser("FDISK")
		paramsFDISK()
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

func pathParams() {
	switch preAnalisis.GetTipo() {
	case "RUTA":
		parser("RUTA")
	case "CADENA":
		parser("CADENA")
	default:
		mensajePanic("(RUTA | \"RUTA\")")
	}
}

func paramsMKDISK() {
	sizeT := false
	pathT := false
	nameT := false
	unitT := false

	for preAnalisis.GetTipo() != "EOF" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "SIZE":
			if sizeT == true {
				panic(">> ERROR PARAMETRO 'SIZE' DUPLICADO")
			}
			sizeT = true
			parser("SIZE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			tamanioDisco()

		case "PATH":
			if pathT == true {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "NAME":
			if nameT == true {
				panic(">> ERROR PARAMETRO 'NAME' DUPLICADO")
			}
			nameT = true
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")
			parser("SIMBOLO_PUNTO")
			parser("DSK")

		case "UNIT":
			if unitT == true {
				panic(">> ERROR PARAMETRO 'UNIT' DUPLICADO")
			}
			unitT = true
			parser("UNIT")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-SIZE | -PATH | -NAME | -UNIT)"
			panic(err)
		}
	}

	if !sizeT || !pathT || !nameT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-SIZE | -PATH | -NAME)")
	}
}

func paramsFDISK() {
	sizeT := false
	unitT := false
	pathT := false
	typeT := false
	fitT := false
	deleteT := false
	nameT := false
	addT := false

	for preAnalisis.GetTipo() != "EOF" {
		parser("SIMBOLO_MENOS")
		switch preAnalisis.GetTipo() {
		case "SIZE":
			if sizeT == true {
				panic(">> ERROR PARAMETRO 'SIZE' DUPLICADO")
			}
			sizeT = true
			parser("SIZE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			tamanioDisco()

		case "PATH":
			if pathT == true {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "UNIT":
			if unitT == true {
				panic(">> ERROR PARAMETRO 'UNIT' DUPLICADO")
			}
			unitT = true
			parser("UNIT")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "TYPE":
			if typeT == true {
				panic(">> ERROR PARAMETRO 'TYPE' DUPLICADO")
			}
			typeT = true
			parser("TYPE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "FIT":
			if fitT == true {
				panic(">> ERROR PARAMETRO 'FIT' DUPLICADO")
			}
			fitT = true
			parser("FIT")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "DELETE":
			if deleteT == true {
				panic(">> ERROR PARAMETRO 'DELTE' DUPLICADO")
			}
			if addT == true {
				panic(">> ERROR PARAMETROS 'DELTE' Y 'ADD INCOMPATIBLES")
			}
			deleteT = true
			parser("DELETE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "NAME":
			if nameT == true {
				panic(">> ERROR PARAMETRO 'NAME' DUPLICADO")
			}
			nameT = true
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			nombresArchivo()

		case "ADD":
			if addT == true {
				panic(">> ERROR PARAMETRO 'DELTE' DUPLICADO")
			}
			if deleteT == true {
				panic(">> ERROR PARAMETROS 'DELTE' Y 'ADD INCOMPATIBLES")
			}
			addT = true
			parser("ADD")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			tamanioDisco()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-SIZE | -PATH | -NAME | -UNIT " +
				"| -TYPE | -FIT | -DELETE | -ADD)"
			panic(err)
		}
	}

	if addT || deleteT {
		if !pathT || !nameT {
			panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-SIZE | -NAME)")
		}

		if fitT || typeT || sizeT || (unitT && deleteT) /*|| (a && d)*/ {
			panic(">> PARAMETROS NO ESPERADOS")
		}
	} else if !sizeT || !pathT || !nameT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-SIZE | -PATH | -NAME)")
	}

}

func tamanioDisco() {
	switch preAnalisis.GetTipo() {
	case "ENTERO":
		parser("ENTERO")
	case "DECIMAL":
		parser("DECIMAL")
	default:
		err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (NUMERO | DECIMAL)"
		panic(err)
	}
}

func nombresArchivo() {
	switch preAnalisis.GetTipo() {
	case "ID":
		parser("ID")
	case "CADENA":
		parser("CADENA")
	default:
		err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (NOMBRE | \"NOMBRE\")"
		panic(err)
	}
}

func parser(tipo string) {
	if preAnalisis.GetTipo() != "EOF" {
		if preAnalisis.GetTipo() != tipo {
			mensajePanic(tipo)
		}

		index++
		preAnalisis = listaTokens[index]
	} else {
		mensajePanic("PARAMETROS")
	}
}

func mensajePanic(mensaje string) {
	tipo := preAnalisis.GetValor()
	if tipo == "EOF" {
		tipo = "FIN DE INSTRUCCION"
	}
	err := ">> 'ERROR: " + tipo + " SE ESPERABA " + mensaje + "'"
	panic(err)
}
