package grafica

import (
	"Sistema-de-archivos-LWH/disco/mbr"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

// TablaDisco genera el archivo para la grafica
func TablaDisco(masterBootR mbr.MBR) {
	var auxiliar strings.Builder
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
			auxiliar.WriteString("\n\t\t\t<tr><td cellpadding='0' colspan='2'><table color='orange' cellspacing='0'>")

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

func graficar(filename string, data string) {
	crearDot(filename, data)
	compilarDot(filename)
	// abrirGrafico(filename)
}

func crearDot(filename string, data string) {
	err := ioutil.WriteFile(filename+".dot", []byte(data), 0777)
	if err != nil {
		panic(err)
	}
}

func compilarDot(filename string) {
	comando := string("dot -Tpng " + filename + ".dot -o " + filename + ".png")

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
