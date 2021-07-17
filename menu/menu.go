package menu

import (
	"FileSystem-LWH/analisis/lexico"
	"FileSystem-LWH/analisis/sintactico"
	"FileSystem-LWH/util"
	"FileSystem-LWH/util/archivo"
	"fmt"
	"strings"
)

// Interfaz de línea de comandos
func Interfaz() {
	fmt.Println("╔══════════════════════╗")
	fmt.Println("║      Bienvenido      ║")
	fmt.Println("╚══════════════════════╝")
	fmt.Print(">> ")
	str := util.LecturaTeclado()

	for !strings.EqualFold(str, "exit") {
		if strings.EqualFold(str, "pause") {
			str = util.LecturaTeclado()
		} else {
			listaTokens, listaErrores := lexico.Scanner(str)
			if len(listaErrores) > 0 {
				fmt.Println(">> 'La entrada contiene errores lexicos'")
				fmt.Println(">> LISTADO DE ERRORES:", listaErrores)
				fmt.Println(">> LISTADO DE TOKENS:", listaTokens)
				fmt.Println()
			} else if len(listaTokens) > 0 {
				if listaTokens[0].GetTipo() == "EXEC" {
					if len(listaTokens) == 6 && listaTokens[1].GetTipo() == "SIMBOLO_MENOS" &&
						listaTokens[2].GetTipo() == "PATH" && listaTokens[3].GetTipo() == "SIMBOLO_MENOS" &&
						listaTokens[4].GetTipo() == "SIMBOLO_MAYOR" && (listaTokens[5].GetTipo() == "RUTA" ||
						listaTokens[5].GetTipo() == "CADENA") && strings.Contains(listaTokens[5].GetValor(), ".mia") {
						archivo.Leer(listaTokens[5].GetValor())
					} else {
						fmt.Println(">> 'ERROR DE INSTRUCCION'")
						fmt.Println()
					}
				} else {
					// INICIO ANALISIS SINTACTICO
					sintactico.Analizar(listaTokens)
				}
			}
		}
		fmt.Print(">> ")
		str = util.LecturaTeclado()
	}
}
