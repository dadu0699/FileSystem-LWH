package mbr

import (
	"Sistema-de-archivos-LWH/disco/particion"
	"math/rand"
	"time"
)

// MBR modelo de la estructura
type MBR struct {
	Tamanio       int64
	FechaCreacion [19]byte
	DiskSignature int64
	Particiones   [4]particion.Particion
}

// Inicializar recibe un puntero MBR para ser modificado.
func (m *MBR) Inicializar(Tamanio int64) {
	m.Tamanio = Tamanio
	fecha := time.Now().Format("01-02-2006 15:04:05")
	copy(m.FechaCreacion[:], fecha)
	m.DiskSignature = rand.Int63n(1000)
}

// SetTamanio recibe un puntero MBR para ser modificado.
func (m *MBR) SetTamanio(Tamanio int64) {
	m.Tamanio = Tamanio
}

// GetTamanio recibe una copia de MBR ya que no necesita modificarlo.
func (m MBR) GetTamanio() int64 {
	return m.Tamanio
}

// GetFecha recibe una copia de MBR ya que no necesita modificarlo.
func (m MBR) GetFecha() string {
	return string(m.FechaCreacion[:])
}

// SetDiskSignature asigna el número random, que identificará de forma única a cada disco
func (m *MBR) SetDiskSignature(DiskSignature int64) {
	m.DiskSignature = DiskSignature
}

// GetDiskSignature devuelve el número random
func (m MBR) GetDiskSignature() int64 {
	return m.DiskSignature
}

// SetParticion recive la posicion y la particion
func (m *MBR) SetParticion(posicion int, particion particion.Particion) {
	m.Particiones[posicion] = particion
}

// GetParticion retorna una particion especifica
func (m MBR) GetParticion(posicion int) particion.Particion {
	return m.Particiones[posicion]
}

// GetParticionPuntero retorna una particion especifica
func (m *MBR) GetParticionPuntero(posicion int) *particion.Particion {
	return &m.Particiones[posicion]
}

// GetParticiones retorna el arreglo de Particiones
func (m MBR) GetParticiones() [4]particion.Particion {
	return m.Particiones
}
