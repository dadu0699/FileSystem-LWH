package escritura

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

// CrearDisco crea el archivo binario
func CrearDisco(tamanio int64, ruta string, nombre string, unidad string) {
	archivo, err := os.Create("./prueba.dsk")
	//crearDirectorio(ruta)

	// Creacion del archivo
	// archivo, err := os.Create(ruta + nombre)
	defer func() {
		archivo.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if err != nil {
		panic(">> 'Error al crear disco'\n")
	}

	switch strings.ToLower(unidad) {
	case "k":
		tamanio *= 1024
	case "m":
		fallthrough
	default:
		tamanio *= (1024 * 1024)
	}

	var inicio int8 = 0
	s := &inicio
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, s)
	escribirEnDisco(archivo, binario.Bytes())

	archivo.Seek(tamanio, 0)
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s)
	escribirEnDisco(archivo, binario2.Bytes())
}

func escribirEnDisco(archivo *os.File, bytes []byte) {
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
