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
		parser("IDN")
		parser("ENTERO")
		parser("SIMBOLO_MENOS")
		parser("SIMBOLO_MAYOR")
		parser("ID")
		listadoIDN()
	case "MKFS":
		parser("MKFS")
		paramsMKFS()
	case "LOGIN":
		parser("LOGIN")
		paramsLOGIN()
	case "LOGOUT":
		parser("LOGOUT")
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
			identificadores()

		case "ADD":
			if addT == true {
				panic(">> ERROR PARAMETRO 'ADD' DUPLICADO")
			}
			if deleteT == true {
				panic(">> ERROR PARAMETROS 'ADD' Y 'DELETE INCOMPATIBLES")
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

	for preAnalisis.GetTipo() != "EOF" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
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

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-SIZE | -PATH | -NAME | -UNIT)"
			panic(err)
		}
	}

	if (pathT || nameT) && (!pathT || !nameT) {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-PATH | -NAME)")
	}
}

func listadoIDN() {
	switch preAnalisis.GetTipo() {
	case "-":
		parser("-")
		parser("IDN")
		parser("ENTERO")
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

	for preAnalisis.GetTipo() != "EOF" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT == true {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "TYPE":
			if typeT == true {
				panic(">> 'ERROR PARAMETRO 'TYPE' DUPLICADO")
			}
			typeT = true
			parser("TYPE")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "ADD":
			if addT == true {
				panic(">> ERROR PARAMETRO 'ADD' DUPLICADO")
			}
			addT = true
			parser("ADD")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			tamanioDisco()

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
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -TYPE | -ADD | -UNIT)"
			panic(err)
		}
	}

	if !idT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID)")
	}
}

func paramsLOGIN() {
	usrT := false
	pwdT := false
	idT := false

	for preAnalisis.GetTipo() != "EOF" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "USR":
			if usrT == true {
				panic(">> ERROR PARAMETRO 'USR' DUPLICADO")
			}
			usrT = true
			parser("USR")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "PWD":
			if pwdT == true {
				panic(">> 'ERROR PARAMETRO 'PWD' DUPLICADO")
			}
			pwdT = true
			parser("PWD")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "IDN":
			if idT == true {
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

	if !usrT || !pwdT || !idT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-USR | -PWD | -ID)")
	}
}

func paramsMKGRP() {
	idT := false
	nameT := false

	for preAnalisis.GetTipo() != "EOF" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT == true {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
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
			identificadores()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -NAME)"
			panic(err)
		}
	}

	if !nameT || !idT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-ID | -NAME)")
	}
}

func paramsRMGRP() {
	idT := false
	nameT := false

	for preAnalisis.GetTipo() != "EOF" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "IDN":
			if idT == true {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
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
			identificadores()

		default:
			err := ">> 'ERROR: " + preAnalisis.GetValor() + " SE ESPERABA (-ID | -NAME)"
			panic(err)
		}
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

	for preAnalisis.GetTipo() != "EOF" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "USR":
			if usrT == true {
				panic(">> ERROR PARAMETRO 'USR' DUPLICADO")
			}
			usrT = true
			parser("USR")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "PWD":
			if pwdT == true {
				panic(">> 'ERROR PARAMETRO 'PWD' DUPLICADO")
			}
			pwdT = true
			parser("PWD")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "IDN":
			if idT == true {
				panic(">> ERROR PARAMETRO 'ID' DUPLICADO")
			}
			idT = true
			parser("IDN")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			parser("ID")

		case "GRP":
			if grpT == true {
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

	if !usrT || !pwdT || !idT || grpT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-USR | -PWD | -ID | -GRP)")
	}
}

func paramsRMUSR() {
	usrT := false
	idT := false

	for preAnalisis.GetTipo() != "EOF" {
		parser("SIMBOLO_MENOS")

		switch preAnalisis.GetTipo() {
		case "USR":
			if usrT == true {
				panic(">> ERROR PARAMETRO 'USR' DUPLICADO")
			}
			usrT = true
			parser("USR")
			parser("SIMBOLO_MENOS")
			parser("SIMBOLO_MAYOR")
			identificadores()

		case "IDN":
			if idT == true {
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

	if !usrT || !idT {
		panic(">> 'ERROR: SE ESPERABAN PARAMETROS OBLIGATORIOS (-USR | -ID)")
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
