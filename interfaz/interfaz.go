package interfaz

import (
	"Sistema-de-archivos-LWH/analisis/lexico"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// MenuPrincipal diseño del menu de consola
func MenuPrincipal() {
	fmt.Println("-------- Bienvenido --------")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		str, _ := reader.ReadString('\n')
		str = strings.Replace(str, "\n", "", -1)
		str = strings.TrimSpace(strings.ToLower(str))

		if str != "exit" {
			lexico.Inicializar()
			lexico.Scanner(str)
			if len(lexico.ListaErrores()) > 0 {
				fmt.Println("--- La entrada contiene errores")
			} else {
				fmt.Println(lexico.ListaTokens())
			}
		} else {
			break
		}
	}
}
