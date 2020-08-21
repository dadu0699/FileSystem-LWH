package archivo

import (
	"Sistema-de-archivos-LWH/analisis/lexico"
	"Sistema-de-archivos-LWH/analisis/sintactico"
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
		linea := scanner.Text()
		if linea != "" {
			fmt.Println(">>", linea)
			listaTokens, listaErrores := lexico.Scanner(linea)
			if len(listaErrores) > 0 {
				fmt.Println(">> 'La entrada contiene errores lexicos'")
				fmt.Println(">> LISTADO DE ERRORES:", listaErrores)
				fmt.Println(">> LISTADO DE TOKENS:", listaTokens)
				fmt.Println()
			} else if len(listaTokens) > 0 {
				if listaTokens[0].GetTipo() == "EXEC" {
					if len(listaTokens) == 4 && listaTokens[1].GetTipo() == "-PATH" &&
						listaTokens[2].GetTipo() == "ASIGNACION" &&
						(listaTokens[3].GetTipo() == "RUTA" ||
							listaTokens[3].GetTipo() == "CADENA") {

						Leer(listaTokens[3].GetValor())
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
