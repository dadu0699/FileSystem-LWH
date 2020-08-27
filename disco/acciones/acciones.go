package acciones

import (
	"Sistema-de-archivos-LWH/disco/mbr"
	"Sistema-de-archivos-LWH/disco/particion"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"sort"
	"strings"
	"unsafe"
)

// Global
var masterBootR mbr.MBR

// CrearDisco crea el archivo binario
func CrearDisco(size int64, path string, name string, unit string) {
	crearDirectorio(path)

	// Creacion del archivo
	file, err := os.Create(path + name)
	defer func() {
		file.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if err != nil {
		panic(err)
	}

	// Asignacion de tamaño especificado por la unidad
	switch strings.ToLower(unit) {
	case "k":
		size *= 1024
	case "m":
		fallthrough
	default:
		size *= (1024 * 1024)
	}

	// Contenido inicial
	var finalCharacter [2]byte
	copy(finalCharacter[:], "\\0")

	// Tamaño de disco
	file.Seek(size, 0) // Posicion Byte final
	var binaryCharacter bytes.Buffer
	binary.Write(&binaryCharacter, binary.BigEndian, &finalCharacter)
	escribirBytes(file, binaryCharacter.Bytes())

	// Inserccion del Master Boot Record
	file.Seek(0, 0) // Posicion Byte inicial
	var masterBootR mbr.MBR
	masterBootR.Inicializar(size)
	var binaryMBR bytes.Buffer
	binary.Write(&binaryMBR, binary.BigEndian, &masterBootR)
	escribirBytes(file, binaryMBR.Bytes())
}

func escribirBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func crearDirectorio(path string) {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}
}

// MontarDisco (lectura) de disco
func MontarDisco(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer func() {
		file.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if err != nil {
		panic(">> 'Error al montar disco'\n")
	}

	var sizeMBR int = int(unsafe.Sizeof(masterBootR))
	data := leerBytes(file, sizeMBR)
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &masterBootR)
	if err != nil {
		panic(err)
	}
}

func leerBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		panic(err)
	}
	return bytes
}

// EliminarDisco remueve el archivo .dsk
func EliminarDisco(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

// CrearParticion crear el struct y lo agrega en el mbr
func CrearParticion(size int64, path string, name string, unit string,
	typeS string, fit string, addT int64, delelteS string) {
	MontarDisco(path)

	// Asignacion de tamaño especificado por la unidad
	switch strings.ToLower(unit) {
	case "b":
		size *= 1
	case "m":
		size *= (1024 * 1024)
	case "k":
		fallthrough
	default:
		size *= 1024
	}

	if (masterBootR.GetTamanio() - int64(unsafe.Sizeof(masterBootR))) >= size {
		posInicial := int64(unsafe.Sizeof(masterBootR))
		particionesLibres := make([]particion.Particion, 0)

		for i := 0; i < 4; i++ {
			particionAuxiliar := masterBootR.GetParticion(i)
			if particionAuxiliar.GetEstado() == byte(0) {
				anterior := particionActivaAnterior(i)
				siguiente := particionActivaSiguiente(i)
				if anterior == -1 && siguiente == -1 {
					particionAuxiliar.SetTamanio(masterBootR.Tamanio - posInicial)
				} else if anterior == -1 && siguiente != -1 {
					particionAuxiliar.SetTamanio(masterBootR.GetParticion(siguiente).GetInicio() - posInicial)
				} else if anterior != -1 && siguiente == -1 {
					particionAuxiliar.SetTamanio(masterBootR.Tamanio -
						(masterBootR.GetParticion(anterior).GetInicio() +
							masterBootR.GetParticion(anterior).GetTamanio()))
				}
				particionesLibres = append(particionesLibres, particionAuxiliar)
			}
		}

		sort.SliceStable(particionesLibres, func(i, j int) bool {
			return particionesLibres[i].Tamanio < particionesLibres[j].Tamanio
		})
		fmt.Println("By age,name:", particionesLibres)
	} else {
		fmt.Println(">> LA PARTICION ES MAS GRANDE QUE EL DISCO")
	}
}

func particionActivaAnterior(posicion int) int {
	for i := posicion; i >= 0; i-- {
		if masterBootR.GetParticion(i).GetEstado() == byte(1) {
			return i
		}
	}
	return -1
}

func particionActivaSiguiente(posicion int) int {
	for i := posicion; i < 4; i++ {
		if masterBootR.GetParticion(i).GetEstado() == byte(1) {
			return i
		}
	}
	return -1
}
