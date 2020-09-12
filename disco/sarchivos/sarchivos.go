package sarchivos

import (
	"Sistema-de-archivos-LWH/disco/acciones"
	"fmt"
	"sort"
	"strconv"
)

// Mount Estructura de las particiones montadas
type Mount struct {
	Ruta   string
	Nombre string
	ID     string
	Numero int
	IDent  string
}

var particionesMontadas []Mount
var post int = 0
var abc = []string{"a", "b", "c", "d", "e", "g", "h", "i", "j", "k", "l", "m"}

// Montar agrega al arreglo la particion
func Montar(path string, name string) {
	acciones.LeerMBR(path)
	acciones.BuscarParticionCreada(name, path)

	sort.SliceStable(particionesMontadas, func(i, j int) bool {
		return particionesMontadas[i].Numero > particionesMontadas[j].Numero
	})

	if buscarParticion(name) {
		panic(">> LA PARTICION FUE MONTADA PREVIAMENTE")
	}

	for _, partition := range particionesMontadas {
		if partition.Ruta == path {
			num := partition.Numero + 1
			particionesMontadas = append(particionesMontadas, Mount{
				Ruta:   path,
				Nombre: name,
				ID:     "vd" + partition.IDent + strconv.Itoa(num),
				Numero: partition.Numero + 1,
				IDent:  partition.IDent,
			})
			return
		}
	}

	particionesMontadas = append(particionesMontadas, Mount{
		Ruta:   path,
		Nombre: name,
		ID:     "vd" + abc[post] + "1",
		Numero: 1,
		IDent:  abc[post],
	})
	post++

	sort.SliceStable(particionesMontadas, func(i, j int) bool {
		return particionesMontadas[i].ID < particionesMontadas[j].ID
	})
}

// MostrarMount imprime todas las particiones montadas en memoria
func MostrarMount() {
	for _, partition := range particionesMontadas {
		fmt.Println(">> -id->"+partition.ID, "-path->\""+partition.Ruta+
			"\" -name->\""+partition.Nombre+"\"")
	}
}

// Desmontar quita la particion de memoria
func Desmontar(name string) {
	sort.SliceStable(particionesMontadas, func(i, j int) bool {
		return particionesMontadas[i].Numero > particionesMontadas[j].Numero
	})

	for _, partition := range particionesMontadas {
		if partition.ID == name {
			particionesMontadas = removeIt(partition, particionesMontadas)
			return
		}
	}
	panic(">> NO SE ENCONTRO LA PARTICION")
}

func removeIt(ss Mount, ssSlice []Mount) []Mount {
	for idx, v := range ssSlice {
		if v == ss {
			return append(ssSlice[0:idx], ssSlice[idx+1:]...)
		}
	}
	return ssSlice
}

func buscarParticion(name string) bool {
	for _, partition := range particionesMontadas {
		if partition.Nombre == name {
			return true
		}
	}
	return false
}

// ObtenerPath regresa la path de una particion
func ObtenerPath(id string) string {
	for _, partition := range particionesMontadas {
		if partition.ID == id {
			return partition.Ruta
		}
	}
	panic(">> NO SE ENCONTRO LA PARTICION")
}
