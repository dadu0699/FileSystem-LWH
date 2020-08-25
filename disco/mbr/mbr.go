package mbr

import (
	"Sistema-de-archivos-LWH/disco/particion"
	"math/rand"
	"time"
)

// MBR modelo de la estructura
type MBR struct {
	tamanio       int64
	fechaCreacion time.Time
	diskSignature int64
	particiones   [4]particion.Particion
}

// Inicializar recibe un puntero MBR para ser modificado.
func (m *MBR) Inicializar(tamanio int64) {
	m.tamanio = tamanio
	m.fechaCreacion = time.Now()
	m.diskSignature = rand.Int63n(1000)
}

// SetTamanio recibe un puntero MBR para ser modificado.
func (m *MBR) SetTamanio(tamanio int64) {
	m.tamanio = tamanio
}

// GetTamanio recibe una copia de MBR ya que no necesita modificarlo.
func (m MBR) GetTamanio() int64 {
	return m.tamanio
}

// GetFecha recibe una copia de MBR ya que no necesita modificarlo.
func (m MBR) GetFecha() string {
	return m.fechaCreacion.Format("01-02-2006 15:04:05")
}

// SetDiskSignature asigna el número random, que identificará de forma única a cada disco
func (m *MBR) SetDiskSignature(diskSignature int64) {
	m.diskSignature = diskSignature
}

// GetDiskSignature devuelve el número random
func (m MBR) GetDiskSignature() int64 {
	return m.diskSignature
}

// SetParticion recive la posicion y la particion
func (m *MBR) SetParticion(posicion int, particion particion.Particion) {
	m.particiones[posicion] = particion
}

// GetParticion retorna una particion especifica
func (m MBR) GetParticion(posicion int) particion.Particion {
	return m.particiones[posicion]
}

// GetParticiones retorna el arreglo de particiones
func (m MBR) GetParticiones() [4]particion.Particion {
	return m.particiones
}
