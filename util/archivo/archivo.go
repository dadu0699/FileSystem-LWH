package archivo

import (
	"FileSystem-LWH/analisis/lexico"
	"FileSystem-LWH/analisis/sintactico"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Leer abre y recorre el archivo
func Leer(ruta string) {
	ruta = strings.ReplaceAll(ruta, "\"", "")
	archivo, err := os.Open(ruta)
	defer func() {
		archivo.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if err != nil {
		panic(">> 'El fichero o directorio no existe'\n")
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		if linea := scanner.Text(); linea != "" {

			for strings.Contains(linea, "\\*") {
				linea = strings.ReplaceAll(linea, "\\*", "")
				scanner.Scan()
				linea += scanner.Text()
			}

			fmt.Println(">>", linea)
			listaTokens, listaErrores := lexico.Scanner(linea)
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
						listaTokens[5].GetTipo() == "CADENA") {
						Leer(listaTokens[5].GetValor())
					} else {
						panic(">> 'ERROR DE INSTRUCCION'\n")
					}
				} else {
					// INICIO ANALISIS SINTACTICO
					sintactico.Analizar(listaTokens)
				}
			}
		}
	}
}
