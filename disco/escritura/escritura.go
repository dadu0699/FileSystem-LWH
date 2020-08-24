package escritura

import (
	"fmt"
	"os"
)

// CrearDisco crea el archivo binario
func CrearDisco(tamanio int64, ruta string, nombre string, unidad string) {
	crearDirectorio(ruta)

	// Creacion del archivo
	archivo, err := os.Create(ruta + nombre)
	defer func() {
		archivo.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if err != nil {
		panic(">> 'Error al crear disco'\n")
	}

}

func escribirDisco(archivo *os.File, bytes []byte) {
	_, err := archivo.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func crearDirectorio(ruta string) {
	err := os.MkdirAll(ruta, 0777)
	if err != nil {
		panic(err)
	}
}
