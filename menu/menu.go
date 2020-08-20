package menu

import (
	"Sistema-de-archivos-LWH/analisis/lexico"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Interfaz de línea de comandos
func Interfaz() {
	fmt.Println("╔══════════════════════╗")
	fmt.Println("║      Bienvenido      ║")
	fmt.Println("╚══════════════════════╝")
	leerArchivo("./entradaPrueba.mias")
	fmt.Print(">> ")
	str := lecturaTeclado()

	for !strings.EqualFold(str, "exit") {
		if strings.EqualFold(str, "pause") {
			str = lecturaTeclado()
		} else {
			lexico.Inicializar()
			lexico.Scanner(str)
			if len(lexico.ListaErrores()) > 0 {
				fmt.Println(">> 'La entrada contiene errores'")
			} else {
				fmt.Println(">>", lexico.ListaTokens())
			}
		}
		fmt.Print(">> ")
		str = lecturaTeclado()
	}
}

func lecturaTeclado() string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.Replace(str, "\n", "", -1)
	return strings.TrimSpace(str)
}

func leerArchivo(ruta string) {
	archivo, err := os.Open(ruta)
	defer archivo.Close()

	if err != nil {
		fmt.Println(">> 'Error al leer archivo'")
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		if linea != "" {
			fmt.Println(">>", linea)
			lexico.Inicializar()
			lexico.Scanner(linea)
			fmt.Println(">>", lexico.ListaTokens(), "\n----------------------------------------------")
		}
	}
}
