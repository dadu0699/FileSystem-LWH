package acciones

import (
	"Sistema-de-archivos-LWH/disco/ebr"
	"Sistema-de-archivos-LWH/disco/mbr"
	"Sistema-de-archivos-LWH/disco/particion"
	"Sistema-de-archivos-LWH/util"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"unsafe"
)

// Global
var masterBootR mbr.MBR
var ebrR ebr.EBR

// CrearDisco crea el archivo binario
func CrearDisco(size int64, path string, name string, unit string) {
	crearDirectorio(path)

	// Creacion del archivo
	file, err := os.Create(path + name)
	defer func() {
		file.Close()
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

	var masterBoot mbr.MBR
	// Inserccion del Master Boot Record
	file.Seek(0, 0) // Posicion Byte inicial
	masterBoot.Inicializar(size)
	var binaryMBR bytes.Buffer
	binary.Write(&binaryMBR, binary.BigEndian, &masterBoot)
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

// LeerMBR Publico
func LeerMBR(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer func() {
		file.Close()
	}()
	if err != nil {
		panic(">> 'ERROR, NO SE PUDO ENCONTRAR EL ARCHIVO DEL DISCO'")
	}

	sizeMBR := int(unsafe.Sizeof(masterBootR))
	data := leerBytes(file, sizeMBR)
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &masterBootR)
	if err != nil {
		panic(err)
	}
}

func leerEBR(path string, start int64) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer func() {
		file.Close()
	}()
	if err != nil {
		panic(">> 'ERROR, NO SE PUDO ENCONTRAR EL ARCHIVO DEL DISCO'\n")
	}

	file.Seek(0, 0)
	file.Seek(start, 0)
	sizeEBR := int(unsafe.Sizeof(ebrR))
	data := leerBytes(file, sizeEBR)
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &ebrR)
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
	typeS string, fit string) {
	LeerMBR(path)
	buscarParticion(name)

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

	// Asignacion de fit
	switch strings.ToLower(fit) {
	case "bf":
		fit = "B"
	case "ff":
		fit = "F"
	case "wf":
		fallthrough
	default:
		fit = "W"
	}

	// Asignacion tipo de particion
	switch strings.ToLower(typeS) {
	case "e":
		typeS = "E"
	case "l":
		typeS = "L"
	case "p":
		fallthrough
	default:
		typeS = "P"
	}

	if (masterBootR.GetTamanio() - int64(unsafe.Sizeof(masterBootR))) >= size {
		if typeS == "P" || typeS == "E" {
			espaciosLibres := buscarEspacioLibre()

			if len(espaciosLibres) == 0 {
				panic(">> NO HAY ESPACIO DISPONIBLE EN EL DISCO")
			}

			/*
				switch fit {
				case "B":
					// Ordenamiento de la lista de particiones libres de menor a mayor
					sort.SliceStable(particionesLibres, func(i, j int) bool {
						return particionesLibres[i].Tamanio > particionesLibres[j].Tamanio
					})

				case "W":
					// Ordenamiento de la lista de particiones libres de mayor a menor
					sort.SliceStable(particionesLibres, func(i, j int) bool {
						return particionesLibres[i].Tamanio < particionesLibres[j].Tamanio
					})
				}
			*/

			if typeS == "E" {
				for _, partition := range masterBootR.GetParticiones() {
					if string(partition.GetNombre()) != "" && string(partition.GetTipo()) == "E" {
						panic(">> YA EXISTE UNA PARTICION EXTENDIDA")
					}
				}
			}

			if masterBootR.Particiones[3].Estado == byte(1) {
				panic(">> YA EXISTEN 4 PARTICIONES EN EL DISCO")
			}

			/*
				for i, particion := range espaciosLibres {
					fmt.Println("--------------", i)
					fmt.Println("INCIO", particion.GetInicio())
					fmt.Println("TAMAÑO", particion.GetTamanio())
				}
			*/

			for i := 0; i < 4; i++ {
				if masterBootR.Particiones[i].Estado == byte(0) {
					for _, particion := range espaciosLibres {
						if (particion.Tamanio) >= size {
							masterBootR.Particiones[i].Estado = 1
							masterBootR.Particiones[i].Tipo = byte(typeS[0])
							masterBootR.Particiones[i].Fit = byte(fit[0])
							masterBootR.Particiones[i].Inicio = particion.Inicio
							masterBootR.Particiones[i].Tamanio = size
							copy(masterBootR.Particiones[i].Nombre[:], name)
							for j := len(name); j < 16; j++ {
								masterBootR.Particiones[i].Nombre[j] = byte(" "[0])
							}
							ordernarParticiones()
							actualizarMBR(path)
							return
						}
					}
				}
			}

			panic(">> LA PARTICION ES MUY GRANDE")
		} else if typeS == "L" {
			for _, partition := range masterBootR.GetParticiones() {
				if string(partition.GetNombre()) != "" && partition.GetTipo() == byte("E"[0]) {
					leerEBR(path, partition.Inicio)

					if ebrR.GetNombre() == "" {
						if int64(unsafe.Sizeof(ebrR))+size <= partition.Tamanio {
							ebrR.Inicializar(byte(fit[0]), partition.Inicio+int64(unsafe.Sizeof(ebrR))+1, size, 0, name)
							actualizarEBR(path, partition.Inicio)
						} else {
							panic(">> TAMAÑO DE PARTICION LOGICA MUY GRANDE")
						}
					} else {
						for ebrR.Siguiente != 0 {
							if ebrR.GetNombre() == name { //TODO NOMBRES SENSITIVE CASE?
								panic(">> YA EXISTE UNA PARTICION LOGICA CON ESE NOMBRE")
							}
							leerEBR(path, ebrR.Siguiente)
						}
						if (ebrR.Inicio+ebrR.Tamanio < partition.Tamanio) &&
							ebrR.Inicio+ebrR.Tamanio+int64(unsafe.Sizeof(ebrR))+size <=
								partition.Tamanio {
							ebrR.Siguiente = ebrR.Inicio + ebrR.Tamanio + 1
							actualizarEBR(path, ebrR.Inicio-int64(unsafe.Sizeof(ebrR))-1)

							var nuevoEbr ebr.EBR
							nuevoEbr.Inicializar(byte(fit[0]), ebrR.Siguiente+int64(unsafe.Sizeof(ebrR))+1, size, 0, name)
							ebrR = nuevoEbr
							actualizarEBR(path, ebrR.Inicio-int64(unsafe.Sizeof(ebrR))-1)
						} else {
							panic(">> TAMAÑO DE PARTICION MUY GRANDE")
						}
					}
					return
				}
			}
			panic(">> NO SE ENCONTRO UNA PARTICION EXTENDIDA")
		}
	} else {
		panic(">> LA PARTICION ES MAS GRANDE QUE EL DISCO")
	}
}

func ordernarParticiones() {
	// ORDEN de inicio menor a mayor
	/*
		sort.SliceStable(masterBootR.Particiones, func(i, j int) bool {
			return masterBootR.Particiones[i].Inicio > masterBootR.Particiones[j].Inicio
		})
	*/

	for i, particion := range masterBootR.Particiones {
		if particion.GetEstado() == byte(0) {
			masterBootR.Particiones[i].Inicio = masterBootR.Tamanio
		}
	}

	n := 4
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if masterBootR.Particiones[i-1].Inicio > masterBootR.Particiones[i].Inicio {
				masterBootR.Particiones[i], masterBootR.Particiones[i-1] = masterBootR.Particiones[i-1], masterBootR.Particiones[i]
				swapped = true
			}
		}
	}
}

func buscarEspacioLibre() []particion.Particion {
	inicioParticiones := int64(unsafe.Sizeof(masterBootR)) + 1

	var espaciosLibres []particion.Particion
	var particionAux particion.Particion

	ordernarParticiones()

	for i, particion := range masterBootR.Particiones {
		if i == 0 {
			if particion.GetEstado() == byte(1) {
				particionAux.Inicio = inicioParticiones
				particionAux.Tamanio = (particion.GetInicio() - 1) - particionAux.GetInicio()
				espaciosLibres = append(espaciosLibres, particionAux)
			} else if particion.GetEstado() == byte(0) {
				particionAux.Inicio = inicioParticiones
				particionAux.Tamanio = masterBootR.GetTamanio() - particionAux.GetInicio()
				espaciosLibres = append(espaciosLibres, particionAux)
				break
			}
		} else if i == 3 {
			if particion.GetEstado() == byte(1) {
				particionAux.Inicio = particion.GetInicio() + particion.GetTamanio() + 1
				particionAux.Tamanio = masterBootR.GetTamanio() - particion.GetTamanio()
				espaciosLibres = append(espaciosLibres, particionAux)
			}
		} else if i > 0 && i < 3 {
			anterior := particionActivaAnterior(i - 1)
			siguiente := particionActivaSiguiente(i + 1)

			if particion.GetEstado() == byte(1) {
				if anterior == -1 {
					particionAux.Inicio = inicioParticiones
					particionAux.Tamanio = particion.GetInicio() - 1 - particionAux.GetInicio()
					espaciosLibres = append(espaciosLibres, particionAux)
				} else if anterior != -1 {
					particionAux.Inicio = masterBootR.GetParticion(anterior).Inicio +
						masterBootR.GetParticion(anterior).GetTamanio() + 1
					particionAux.Tamanio = particion.GetInicio() - 1 - particionAux.GetInicio()
					espaciosLibres = append(espaciosLibres, particionAux)
				}

				if siguiente == -1 {
					particionAux.Inicio = particion.GetInicio() + particion.GetTamanio() + 1
					particionAux.Tamanio = masterBootR.GetTamanio() - particionAux.GetInicio()
					espaciosLibres = append(espaciosLibres, particionAux)
				} else if siguiente != -1 {
					particionAux.Inicio = particion.GetInicio() + particion.GetTamanio() + 1
					particionAux.Tamanio = masterBootR.GetParticion(siguiente).GetInicio() - 1 - particionAux.GetInicio()
					espaciosLibres = append(espaciosLibres, particionAux)
				}
			} else if particion.GetEstado() == byte(0) {
				if anterior == -1 && siguiente == -1 {
					particionAux.Inicio = inicioParticiones
					particionAux.Tamanio = masterBootR.GetTamanio() - inicioParticiones
					espaciosLibres = append(espaciosLibres, particionAux)
					break
				} else if anterior == -1 && siguiente != -1 {
					particionAux.Inicio = inicioParticiones
					particionAux.Tamanio = masterBootR.GetParticion(siguiente).GetInicio() - 1 - inicioParticiones
					espaciosLibres = append(espaciosLibres, particionAux)
				} else if anterior != -1 && siguiente == -1 {
					particionAux.Inicio = masterBootR.GetParticion(anterior).GetInicio() +
						masterBootR.GetParticion(anterior).Tamanio + 1
					particionAux.Tamanio = masterBootR.GetTamanio() -
						masterBootR.GetParticion(anterior).GetInicio() -
						masterBootR.GetParticion(anterior).GetTamanio() + 1
					espaciosLibres = append(espaciosLibres, particionAux)
				} else if anterior != -1 && siguiente != -1 {
					particionAux.Inicio = masterBootR.GetParticion(anterior).GetInicio() +
						masterBootR.GetParticion(anterior).GetTamanio() + 1
					particionAux.Tamanio = masterBootR.GetParticion(siguiente).GetInicio() - 1
					espaciosLibres = append(espaciosLibres, particionAux)
				}
			}
		}
	}
	return espaciosLibres
}

func buscarParticion(nombre string) {
	for _, partition := range masterBootR.GetParticiones() {
		if string(partition.GetNombre()) != "" && string(partition.GetNombre()) == nombre { //TODO NOMBRES SENSITIVE CASE?
			panic(">> YA EXISTE UNA PARTICION CON ESE NOMBRE")
		}
	}
}

// BuscarParticionCreada por medio del nombre
func BuscarParticionCreada(nombre string) {
	for _, partition := range masterBootR.GetParticiones() {
		if string(partition.GetNombre()) != "" && string(partition.GetNombre()) == nombre { //TODO NOMBRES SENSITIVE CASE?
			return
		}
	}
	panic(">> PARTICION NO ECONTRADA")
}

func buscarPartLibres() []*particion.Particion {
	posInicial := int64(unsafe.Sizeof(masterBootR))
	// Busqueda de particiones libres o desactivas
	particionesLibres := make([]*particion.Particion, 0)
	for i := 0; i < 4; i++ {
		particionAuxiliar := masterBootR.GetParticionPuntero(i)
		if particionAuxiliar.GetEstado() == byte(0) {
			anterior := particionActivaAnterior(i)
			siguiente := particionActivaSiguiente(i)
			if anterior == -1 && siguiente == -1 {
				particionAuxiliar.SetTamanio(masterBootR.Tamanio - posInicial)
				particionAuxiliar.SetInicio(posInicial + 1) // +1
			} else if anterior == -1 && siguiente != -1 {
				particionAuxiliar.SetTamanio(masterBootR.GetParticion(siguiente).GetInicio() - posInicial)
				particionAuxiliar.SetInicio(posInicial + 1) // +1
			} else if anterior != -1 && siguiente == -1 {
				particionAuxiliar.SetTamanio(masterBootR.Tamanio -
					(masterBootR.GetParticion(anterior).GetInicio() +
						masterBootR.GetParticion(anterior).GetTamanio()))
				particionAuxiliar.SetInicio(masterBootR.GetParticion(anterior).GetInicio() +
					masterBootR.GetParticion(anterior).GetTamanio() + 1) // +1
			} else if anterior != -1 && siguiente != -1 {
				particionAuxiliar.SetInicio(masterBootR.GetParticion(anterior).GetInicio() +
					masterBootR.GetParticion(anterior).GetTamanio() + 1)
				particionAuxiliar.SetTamanio(masterBootR.GetParticion(siguiente).GetInicio() -
					(masterBootR.GetParticion(anterior).GetInicio() +
						masterBootR.GetParticion(anterior).GetTamanio())) // VERIFICAR QUE el tamaño se correcto
			}
			particionesLibres = append(particionesLibres, particionAuxiliar)
		}
	}
	return particionesLibres
}

func actualizarMBR(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer func() {
		file.Close()
	}()
	if err != nil {
		panic(">> 'Error al montar disco'\n")
	}

	// Inserccion del Master Boot Record
	file.Seek(0, 0) // Posicion Byte inicial
	var binaryMBR bytes.Buffer
	binary.Write(&binaryMBR, binary.BigEndian, &masterBootR)
	escribirBytes(file, binaryMBR.Bytes())
}

func actualizarEBR(path string, pos int64) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer func() {
		file.Close()
	}()
	if err != nil {
		panic(">> 'Error al montar disco'\n")
	}

	file.Seek(pos, 0) // Posicion Byte inicial
	var binaryEBR bytes.Buffer
	binary.Write(&binaryEBR, binary.BigEndian, &ebrR)
	escribirBytes(file, binaryEBR.Bytes())
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

// EliminarParticion realiza el formateo y eliminacion de una particion
func EliminarParticion(path string, name string, deleteP string) {
	LeerMBR(path)
	for i := 0; i < 4; i++ {
		if strings.EqualFold(masterBootR.GetParticion(i).GetNombre(), name) {
			masterBootR.Particiones[i].Estado = byte(0)
			masterBootR.Particiones[i].Tipo = byte(0)
			masterBootR.Particiones[i].Fit = byte(0)
			copy(masterBootR.Particiones[i].Nombre[:], " ")
			for j := len(" "); j < 16; j++ {
				masterBootR.Particiones[i].Nombre[j] = byte(" "[0])
			}

			if strings.EqualFold(deleteP, "FULL") {
				escribirCeros(path, masterBootR.Particiones[i].Inicio, masterBootR.Particiones[i].Tamanio)
			}

			inicio := int64(0)
			fin := int64(0)

			/*
				posAnterior := particionActivaAnterior(i)
				posSig := particionActivaSiguiente(i)
					if posAnterior != -1 {
						inicio = masterBootR.Particiones[posAnterior].Inicio +
							masterBootR.Particiones[posAnterior].Tamanio + 1
					}

					if posSig != -1 {
						fin = masterBootR.Particiones[posSig].Inicio - 1
					}
			*/
			masterBootR.Particiones[i].Inicio = inicio
			masterBootR.Particiones[i].Tamanio = fin
			ordernarParticiones()
			actualizarMBR(path)
			fmt.Println(">> PARTICION ELIMINADA CORRECTAMENTE")
			util.LecturaTeclado()
			return
		} else if masterBootR.GetParticion(i).GetTipo() == byte("E"[0]) {
			leerEBR(path, masterBootR.GetParticion(i).GetInicio()) //TODO VERIFICAR PARTICIONES LOGICAS
		}
	}
	panic(">> LA PARTICION NO FUE ENCONTRADA")
}

func escribirCeros(path string, inicio int64, fin int64) {
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer func() {
		file.Close()
	}()
	if err != nil {
		panic(">> 'Error al montar disco'\n")
	}

	for i := inicio; i <= (fin + inicio); i++ {
		var finalCharacter [2]byte
		copy(finalCharacter[:], "0")
		file.Seek(i, 0)
		var binaryCharacter bytes.Buffer
		binary.Write(&binaryCharacter, binary.BigEndian, &finalCharacter)
		escribirBytes(file, binaryCharacter.Bytes())
	}
}

// CambiarTamanio metodo el cual aumenta o reduce el tamanio de una particion
func CambiarTamanio(addT int64, path string, name string, unit string) {
	LeerMBR(path)

	switch strings.ToLower(unit) {
	case "b":
		addT *= 1
	case "m":
		addT *= (1024 * 1024)
	case "k":
		fallthrough
	default:
		addT *= 1024
	}

	for i := 0; i < 4; i++ {
		if strings.EqualFold(masterBootR.GetParticion(i).GetNombre(), name) {
			if addT > 0 {
				posSig := particionActivaSiguiente(i + 1)
				if posSig == -1 {
					if masterBootR.GetParticion(i).GetTamanio() < masterBootR.Tamanio &&
						addT <= masterBootR.Tamanio {
						masterBootR.Particiones[i].Tamanio = masterBootR.Particiones[i].Tamanio + addT
					} else {
						panic(">> EL AUMENTO DE TAMAÑO DE LA PARTICION NO PUEDE SER MAYOR AL TAMAÑO ACTUAL DEL DISCO")
					}
				} else {
					part := masterBootR.GetParticion(i)
					if (part.GetInicio()+part.GetTamanio()+1) < (masterBootR.GetParticion(posSig).GetInicio()-1) &&
						addT <= (masterBootR.GetParticion(posSig).GetInicio()-1) {
						masterBootR.Particiones[i].Tamanio = masterBootR.Particiones[i].Tamanio + addT
					} else {
						panic(">> ERROR, NO SE PUEDE AUMENTAR EL TAMAÑO DE LA PARTICION")
					}
				}
			} else {
				if masterBootR.GetParticion(i).GetTamanio() >= (addT * -1) {
					masterBootR.Particiones[i].Tamanio = masterBootR.Particiones[i].Tamanio + addT
				} else {
					panic(">> LA REDUCCION DE LA PARTICION NO PUEDE SER MAYOR AL TAMAÑO ACTUAL")
				}
			}
			ordernarParticiones()
			actualizarMBR(path)
			return
		}
	}
	panic(">> LA PARTICION NO FUE ENCONTRADA")
}
