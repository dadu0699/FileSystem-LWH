package sintactico

import (
	"Sistema-de-archivos-LWH/analisis/token"
	"Sistema-de-archivos-LWH/ejecucion"
	"Sistema-de-archivos-LWH/util"
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
			util.LecturaTeclado()
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
		parser("PAUSE")
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
		parser("MOUNT")
		paramsMOUNT()
	case "UNMOUNT":
		parser("UNMOUNT")
		parser("SIMBOLO_MENOS")
		parser("ID")
		parser("SIMBOLO_MENOS")
		parser("SIMBOLO_MAYOR")
		parser("ID")
		listadoIDN()
	case "MKFS":
		parser("MKFS")
		paramsMKFS()
	case "LOGIN": // TODO LOGIN PUEDEN VENIR NUMEROS como PWD
		parser("LOGIN")
		paramsLOGIN()
	case "LOGOUT":
		parser("LOGOUT")
		if preAnalisis.GetTipo() == "COMENTARIO" {
			parser("COMENTARIO")
		}
	case "MKGRP":
		parser("MKGRP")
		paramsMKGRP()
	case "RMGRP":
		parser("RMGRP")
		paramsRMGRP()
	case "MKUSR":
		parser("MKUSR")
		paramsMKUSR()
	case "RMUSR":
		parser("RMUSR")
		paramsRMUSR()
	case "CHMOD":
		parser("CHMOD")
		paramsCHMOD()
	case "MKFILE":
		parser("MKFILE")
		paramsMKFILE()
	case "CAT":
		parser("CAT")
		paramsCAT()
	case "RM":
		parser("RM")
		paramsRM()
	case "EDIT":
		parser("EDIT")
		paramsEDIT()
	case "REN":
		parser("REN")
		paramsREN()
	case "MKDIR":
		parser("MKDIR")
		paramsMKDIR()
	case "CP":
		parser("CP")
		paramsCP()
	case "MV":
		parser("MV")
		paramsMV()
	case "FIND":
		parser("FIND")
		paramsFIND()
	case "CHOWN":
		parser("CHOWN")
		paramsCHOWN()
	case "CHGRP":
		parser("CHGRP")
		paramsCHGRP()
	case "LOSS":
		parser("LOSS")
		paramsLOSS()
	case "RECOVERY":
		parser("RECOVERY")
		paramsRECOVERY()
	case "REP":
		parser("REP")
		paramsREP()
	}
}

func pathParams() {
	switch preAnalisis.GetTipo() {
	case "RUTA":
		parser("RUTA")
	case "CADENA":
		parser("CADENA")
	default:
		err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (RUTA | \"RUTA\")"
		panic(err)
	}
}

func paramsMKDISK() {
	sizeT := false
	pathT := false
	nameT := false
	unitT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "SIZE":
			if sizeT {
				panic(">> ERROR PARAMETRO 'SIZE' DUPLICADO")
			}
			sizeT = true
			parser("SIZE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			tamanioDisco()

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "NAME":
			if nameT {
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
			if unitT {
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

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
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

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")
		switch preAnalisis.GetTipo() {
		case "SIZE":
			if sizeT {
				panic(">> ERROR PARAMETRO 'SIZE' DUPLICADO")
			}
			sizeT = true
			parser("SIZE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			tamanioDisco()

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "UNIT":
			if unitT {
				panic(">> ERROR PARAMETRO 'UNIT' DUPLICADO")
			}
			unitT = true
			parser("UNIT")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "TYPE":
			if typeT {
				panic(">> ERROR PARAMETRO 'TYPE' DUPLICADO")
			}
			typeT = true
			parser("TYPE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "FIT":
			if fitT {
				panic(">> ERROR PARAMETRO 'FIT' DUPLICADO")
			}
			fitT = true
			parser("FIT")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "DELETE":
			if deleteT {
				panic(">> ERROR PARAMETRO 'DELTE' DUPLICADO")
			}
			if addT {
				panic(">> ERROR PARAMETROS 'DELTE' Y 'ADD INCOMPATIBLES")
			}
			deleteT = true
			parser("DELETE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "NAME":
			if nameT {
				panic(">> ERROR PARAMETRO 'NAME' DUPLICADO")
			}
			nameT = true
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "ADD":
			if addT {
				panic(">> ERROR PARAMETRO 'ADD' DUPLICADO")
			}
			if deleteT {
				panic(">> ERROR PARAMETROS 'ADD' Y 'DELETE INCOMPATIBLES")
			}
			addT = true
			parser("ADD")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			if preAnalisis.GetTipo() == "SIMBOLO_MENOS" {
				parser("SIMBOLO_MENOS")
			}
			tamanioDisco()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-SIZE | -PATH | -NAME | -UNIT " +
				"| -TYPE | -FIT | -DELETE | -ADD)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
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

func identificadores() {
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

func paramsMOUNT() {
	pathT := false
	nameT := false
	parametros := 0

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "NAME":
			if nameT {
				panic(">> ERROR PARAMETRO 'NAME' DUPLICADO")
			}
			nameT = true
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-PATH | -NAME)"
			panic(err)
		}
		parametros++
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if parametros > 0 {
		if (pathT || nameT) && (!pathT || !nameT) {
			panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-PATH | -NAME)")
		}
	}
}

func listadoIDN() {
	switch preAnalisis.GetTipo() {
	case "-":
		parser("-")
		parser("ID")
		parser("SIMBOLO_MENOS")
		parser("SIMBOLO_MAYOR")
		parser("ID")
		listadoIDN()
	}
}

func paramsMKFS() {
	idT := false
	typeT := false
	addT := false
	unitT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "TYPE":
			if typeT {
				panic(">> 'ERROR PARAMETRO 'TYPE' DUPLICADO")
			}
			typeT = true
			parser("TYPE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "ADD":
			if addT {
				panic(">> ERROR PARAMETRO 'ADD' DUPLICADO")
			}
			addT = true
			parser("ADD")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			if preAnalisis.GetTipo() == "SIMBOLO_MENOS" {
				parser("SIMBOLO_MENOS")
			}
			tamanioDisco()

		case "UNIT":
			if unitT {
				panic(">> ERROR PARAMETRO 'UNIT' DUPLICADO")
			}
			unitT = true
			parser("UNIT")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -TYPE | -ADD | -UNIT)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID)")
	}
}

func paramsLOGIN() {
	usrT := false
	pwdT := false
	idT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "USR":
			if usrT {
				panic(">> ERROR PARAMETRO 'USR' DUPLICADO")
			}
			usrT = true
			parser("USR")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "PWD":
			if pwdT {
				panic(">> 'ERROR PARAMETRO 'PWD' DUPLICADO")
			}
			pwdT = true
			parser("PWD")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-USR | -PWD | -ID)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !usrT || !pwdT || !idT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-USR | -PWD | -ID)")
	}
}

func paramsMKGRP() {
	idT := false
	nameT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "NAME":
			if nameT {
				panic(">> ERROR PARAMETRO 'NAME' DUPLICADO")
			}
			nameT = true
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -NAME)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !nameT || !idT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -NAME)")
	}
}

func paramsRMGRP() {
	idT := false
	nameT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "NAME":
			if nameT {
				panic(">> ERROR PARAMETRO 'NAME' DUPLICADO")
			}
			nameT = true
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -NAME)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !nameT || !idT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -NAME)")
	}
}

func paramsMKUSR() {
	usrT := false
	pwdT := false
	idT := false
	grpT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "USR":
			if usrT {
				panic(">> ERROR PARAMETRO 'USR' DUPLICADO")
			}
			usrT = true
			parser("USR")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "PWD":
			if pwdT {
				panic(">> 'ERROR PARAMETRO 'PWD' DUPLICADO")
			}
			pwdT = true
			parser("PWD")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "GRP":
			if grpT {
				panic(">> ERROR PARAMETRO 'GRP' DUPLICADO")
			}
			grpT = true
			parser("GRP")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-USR | -PWD | -ID | -GRP)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !usrT || !pwdT || !idT || !grpT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-USR | -PWD | -ID | -GRP)")
	}
}

func paramsRMUSR() {
	usrT := false
	idT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "USR":
			if usrT {
				panic(">> ERROR PARAMETRO 'USR' DUPLICADO")
			}
			usrT = true
			parser("USR")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-USR | -ID)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !usrT || !idT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-USR | -ID)")
	}
}

func paramsCHMOD() {
	idT := false
	pathT := false
	ugoT := false
	rT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "UGO":
			if ugoT {
				panic(">> ERROR PARAMETRO 'UGO' DUPLICADO")
			}
			ugoT = true
			parser("UGO")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ENTERO")

		case "R":
			if rT {
				panic(">> ERROR PARAMETRO 'R' DUPLICADO")
			}
			rT = true
			parser("R")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -PATH | -UGO | -R)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT || !pathT || !ugoT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -PATH | -UGO)")
	}
}

func paramsMKFILE() {
	idT := false
	pathT := false
	pT := false
	sizeT := false
	contT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "SIZE":
			if sizeT {
				panic(">> ERROR PARAMETRO 'SIZE' DUPLICADO")
			}
			sizeT = true
			parser("SIZE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			tamanioDisco()

		case "P":
			if pT {
				panic(">> ERROR PARAMETRO 'P' DUPLICADO")
			}
			pT = true
			parser("P")

		case "CONT":
			if contT {
				panic(">> ERROR PARAMETRO 'CONT' DUPLICADO")
			}
			contT = true
			parser("CONT")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("CADENA")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -PATH | -P | -SIZE | -CONT)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT || !pathT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -PATH)")
	}
}

func paramsCAT() {
	idT := false
	fileT := 0

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "ID":
			fileT++
			parser("ID")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("CADENA")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -FILE)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT || fileT <= 0 {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -FILE)")
	}
}

func paramsRM() {
	idT := false
	pathT := false
	rfT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "RF":
			if rfT {
				panic(">> ERROR PARAMETRO 'RF' DUPLICADO")
			}
			rfT = true
			parser("RF")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -PATH | -RF)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT || !pathT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -PATH)")
	}
}

func paramsEDIT() {
	idT := false
	pathT := false
	sizeT := false
	contT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "SIZE":
			if sizeT {
				panic(">> ERROR PARAMETRO 'SIZE' DUPLICADO")
			}
			sizeT = true
			parser("SIZE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			tamanioDisco()

		case "CONT":
			if contT {
				panic(">> ERROR PARAMETRO 'CONT' DUPLICADO")
			}
			contT = true
			parser("CONT")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("CADENA")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -PATH | -SIZE | -CONT)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT || !pathT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -PATH)")
	}
}

func paramsREN() {
	idT := false
	pathT := false
	nameT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "NAME":
			if nameT {
				panic(">> ERROR PARAMETRO 'NAME' DUPLICADO")
			}
			nameT = true
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -PATH | -NAME)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT || !pathT || !nameT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -PATH | -NAME)")
	}
}

func paramsMKDIR() {
	idT := false
	pathT := false
	pT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "P":
			if pT {
				panic(">> ERROR PARAMETRO 'P' DUPLICADO")
			}
			pT = true
			parser("P")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -PATH | -P)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT || !pathT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -PATH)")
	}
}

func paramsCP() {
	idT := false
	pathT := false
	destT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "DEST":
			if destT {
				panic(">> ERROR PARAMETRO 'DEST' DUPLICADO")
			}
			destT = true
			parser("DEST")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -PATH | -DEST)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT || !pathT || !destT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -PATH | -DEST)")
	}
}

func paramsMV() {
	idT := false
	pathT := false
	destT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "DEST":
			if destT {
				panic(">> ERROR PARAMETRO 'DEST' DUPLICADO")
			}
			destT = true
			parser("DEST")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -PATH | -DEST)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT || !pathT || !destT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -PATH | -DEST)")
	}
}

func paramsFIND() {
	idT := false
	pathT := false
	nameT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "NAME":
			if nameT {
				panic(">> ERROR PARAMETRO 'NAME' DUPLICADO")
			}
			nameT = true
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			if preAnalisis.GetTipo() == "SIMBOLO_INTERROGACION" {
				parser("SIMBOLO_INTERROGACION")
			} else if preAnalisis.GetTipo() == "SIMBOLO_ASTERISCO" {
				parser("SIMBOLO_ASTERISCO")
			} else {
				identificadores()
			}

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -PATH | -NAME)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT || !pathT || !nameT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -PATH | -NAME)")
	}
}

func paramsCHOWN() {
	idT := false
	pathT := false
	rT := false
	usrT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "R":
			if rT {
				panic(">> ERROR PARAMETRO 'R' DUPLICADO")
			}
			rT = true
			parser("R")

		case "USR":
			if usrT {
				panic(">> ERROR PARAMETRO 'USR' DUPLICADO")
			}
			usrT = true
			parser("USR")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -PATH | -USR | -R)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT || !pathT || !usrT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -PATH | -USR)")
	}
}

func paramsCHGRP() {
	grpT := false
	usrT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "GRP":
			if grpT {
				panic(">> ERROR PARAMETRO 'GRP' DUPLICADO")
			}
			grpT = true
			parser("GRP")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "USR":
			if usrT {
				panic(">> ERROR PARAMETRO 'USR' DUPLICADO")
			}
			usrT = true
			parser("USR")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -GRP)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !grpT || !usrT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -GRP)")
	}
}

func paramsREP() {
	idT := false
	pathT := false
	nameT := false
	rutaT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "NAME":
			if nameT {
				panic(">> ERROR PARAMETRO 'NAME' DUPLICADO")
			}
			nameT = true
			parser("NAME")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "PATH":
			if pathT {
				panic(">> 'ERROR PARAMETRO 'PATH' DUPLICADO")
			}
			pathT = true
			parser("PATH")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		case "RUTA":
			if rutaT {
				panic(">> ERROR PARAMETRO 'RUTA' DUPLICADO")
			}
			rutaT = true
			parser("RUTA")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			pathParams()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -NAME | -PATH | -RUTA)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !nameT || !idT || !pathT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -NAME| -PATH)")
	}
}

func paramsLOSS() {
	idT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID)")
	}
}

func paramsRECOVERY() {
	idT := false

	for preAnalisis.GetTipo() != "EOF" && preAnalisis.GetTipo() != "COMENTARIO" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID)"
			panic(err)
		}
	}

	if preAnalisis.GetTipo() == "COMENTARIO" {
		parser("COMENTARIO")
	}

	if !idT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID)")
	}
}

func parser(tipo string) {
	if preAnalisis.GetTipo() != "EOF" {
		if preAnalisis.GetTipo() != tipo {
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA " + tipo + "'"
			panic(err)
		}

		index++
		preAnalisis = listaTokens[index]
	} else {
		panic(">> 'ERROR: FIN DE ARCHIVO SE ESPERABAN MAS PARAMETROS'")
	}
}
