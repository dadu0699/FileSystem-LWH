package escritura

import (
	"Sistema-de-archivos-LWH/disco/mbr"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
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

	// Asignacion de tamaño especifico
	switch strings.ToLower(unidad) {
	case "k":
		tamanio *= 1024
	case "m":
		fallthrough
	default:
		tamanio *= (1024 * 1024)
	}

	// Contenido inicial
	var caracterFinal [2]byte
	copy(caracterFinal[:], "\\0")

	// Tamaño de disco
	archivo.Seek(tamanio, 0)
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, &caracterFinal)
	escribirEnDisco(archivo, binario2.Bytes())

	// Inserccion del Master Boot Record
	archivo.Seek(0, 0)
	var masterBootR mbr.MBR
	masterBootR.Inicializar(tamanio)
	var binario3 bytes.Buffer
	binary.Write(&binario3, binary.BigEndian, &masterBootR)
	escribirEnDisco(archivo, binario3.Bytes())
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
