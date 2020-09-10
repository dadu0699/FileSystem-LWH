package grafica

import (
	"Sistema-de-archivos-LWH/disco/ebr"
	"Sistema-de-archivos-LWH/disco/mbr"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unsafe"
)

var masterBootR mbr.MBR
var ebrR ebr.EBR

// TablaDisco genera el archivo para la grafica
func TablaDisco(path string) {
	var auxiliar strings.Builder
	leerMBR(path)

	auxiliar.WriteString("digraph G{")
	auxiliar.WriteString("\n\ttbl [ \n\tshape=plaintext \n\tlabel=<")
	auxiliar.WriteString("\n\t\t<table border='0' cellborder='1' color='black' cellspacing='0'>")
	auxiliar.WriteString("\n\t\t\t<tr><td>Nombre</td><td>Valor</td></tr>")

	auxiliar.WriteString("\n\t\t\t<tr><td>MBR_TAMANIO</td><td>")
	auxiliar.WriteString(strconv.FormatInt(masterBootR.GetTamanio(), 10))
	auxiliar.WriteString("</td></tr>")

	auxiliar.WriteString("\n\t\t\t<tr><td>MBR_FECHA_CREACION</td><td>")
	auxiliar.WriteString(masterBootR.GetFecha())
	auxiliar.WriteString("</td></tr>")

	auxiliar.WriteString("\n\t\t\t<tr><td>MBR_DISK_SIGNATURE</td><td>")
	auxiliar.WriteString(strconv.FormatInt(masterBootR.GetDiskSignature(), 10))
	auxiliar.WriteString("</td></tr>")

	for i, partition := range masterBootR.GetParticiones() {
		if name := string(partition.GetNombre()); name != "" {
			auxiliar.WriteString("\n\t\t\t<tr><td cellpadding='0' colspan='2'>")
			auxiliar.WriteString("<table color='orange' cellspacing='0'>")
			auxiliar.WriteString("\n\t\t\t\t<tr><td>PART_ESTADO_")
			auxiliar.WriteString(strconv.Itoa(i + 1))
			auxiliar.WriteString("</td><td>")
			str := fmt.Sprintf("%d", partition.GetEstado())
			auxiliar.WriteString(str)
			auxiliar.WriteString("</td></tr>")

			auxiliar.WriteString("\n\t\t\t\t<tr><td>PART_TIPO_")
			auxiliar.WriteString(strconv.Itoa(i + 1))
			auxiliar.WriteString("</td><td>")
			auxiliar.WriteString(string(partition.GetTipo()))
			auxiliar.WriteString("</td></tr>")

			auxiliar.WriteString("\n\t\t\t\t<tr><td>PART_FIT_")
			auxiliar.WriteString(strconv.Itoa(i + 1))
			auxiliar.WriteString("</td><td>")
			auxiliar.WriteString(string(partition.GetFit()))
			auxiliar.WriteString("</td></tr>")

			auxiliar.WriteString("\n\t\t\t\t<tr><td>PART_INICIO_")
			auxiliar.WriteString(strconv.Itoa(i + 1))
			auxiliar.WriteString("</td><td>")
			auxiliar.WriteString(strconv.FormatInt(partition.GetInicio(), 10))
			auxiliar.WriteString("</td></tr>")

			auxiliar.WriteString("\n\t\t\t\t<tr><td>PART_TAMANIO_")
			auxiliar.WriteString(strconv.Itoa(i + 1))
			auxiliar.WriteString("</td><td>")
			auxiliar.WriteString(strconv.FormatInt(partition.GetTamanio(), 10))
			auxiliar.WriteString("</td></tr>")

			auxiliar.WriteString("\n\t\t\t\t<tr><td>PART_NOMBRE_")
			auxiliar.WriteString(strconv.Itoa(i + 1))
			auxiliar.WriteString("</td><td>")
			auxiliar.WriteString(name)
			auxiliar.WriteString("</td></tr>")
			auxiliar.WriteString("\n\t\t\t</table></td></tr>")
		}
	}

	auxiliar.WriteString("\n\t\t</table>>];}")
	graficar("informacionDisco", auxiliar.String())
}

// RepDisco estructura del .dsk
func RepDisco(path string) {
	var auxiliar strings.Builder
	leerMBR(path)

	auxiliar.WriteString("digraph G{")
	auxiliar.WriteString("\ntbl [ \nshape=plaintext \nlabel=<")
	auxiliar.WriteString("\n<table> \n<tr><td>MBR</td>")

	for i, particion := range masterBootR.GetParticiones() {
		if i == 0 && particion.GetEstado() == 0 {
			auxiliar.WriteString("<td>LIBRE</td>")
			break

		} else {
			anterior := particionActivaAnterior(i - 1)
			siguiente := particionActivaSiguiente(i + 1)
			espacio := int64(0)

			if particion.GetEstado() == 1 {
				if anterior == -1 {
					espacio = particion.GetInicio() - 1 - int64(unsafe.Sizeof(masterBootR))
				} else if anterior != -1 {
					espacio = particion.GetInicio() - 1 - (masterBootR.Particiones[anterior].GetInicio() +
						masterBootR.Particiones[anterior].GetTamanio())
				}
				if espacio > 0 {
					auxiliar.WriteString("<td>LIBRE</td>")
				}

				if particion.GetTipo() == byte("P"[0]) {
					auxiliar.WriteString("<td>PRIMARIA: " + particion.GetNombre() + "</td>")
				} else {
					auxiliar.WriteString("<td>")
					auxiliar.WriteString("<table border='0' cellborder='1' cellspacing='0'>")
					colspan := 1
					logPart := ""

					///////////////////////
					leerEBR(path, particion.GetInicio())
					if ebrR.GetInicio() != 0 || ebrR.GetSiguiente() != 0 {
						logPart += "<tr>"

						if ebrR.GetInicio() != 0 {
							logPart += "<td>EBR</td>"
							logPart += "<td>LOGICA: " + ebrR.GetNombre() + "</td>"
							colspan += 2
						}

						for ebrR.Siguiente != 0 {
							leerEBR(path, ebrR.Siguiente)
							logPart += "<td>EBR</td>"
							logPart += "<td>LOGICA: " + ebrR.GetNombre() + "</td>"
							colspan += 2
						}
						logPart += "</tr>"
					}
					///////////////////////////

					auxiliar.WriteString("<tr><td colspan='")
					str := strconv.Itoa(colspan)
					auxiliar.WriteString(str)
					auxiliar.WriteString("'>EXTENDIDA: " + particion.GetNombre() + "</td></tr>")
					auxiliar.WriteString(logPart)
					auxiliar.WriteString("</table></td>")
				}

				if siguiente == -1 {
					espacio = masterBootR.GetTamanio() - (particion.GetInicio() + particion.GetTamanio() + 1)
					if espacio > 0 {
						auxiliar.WriteString("<td>LIBRE</td>")
					}
					break
				}
			}
		}
	}

	auxiliar.WriteString("\n</tr></table>>];}")
	graficar("disco", auxiliar.String())
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

func graficar(filename string, data string) {
	crearDot(filename, data)
	compilarDot(filename, "png")
	// abrirGrafico(filename)
}

func crearDot(filename string, data string) {
	err := ioutil.WriteFile(filename+".dot", []byte(data), 0777)
	if err != nil {
		panic(err)
	}
}

func compilarDot(filename string, extension string) {
	comando := string("dot -T" + extension + " " + filename + ".dot -o " + filename + "." + extension)

	args := strings.Split(comando, " ")
	cmd := exec.Command(args[0], args[1:]...)

	_, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
}

func abrirGrafico(filename string) {
	comando := string(filename + ".png")

	args := strings.Split(comando, " ")
	cmd := exec.Command(args[0], args[1:]...)

	_, err := cmd.CombinedOutput()
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

func leerMBR(path string) {
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
