package particion

import (
	"strings"
)

// Particion modelo de la estructura
type Particion struct {
	Estado  byte
	Tipo    byte
	Fit     byte
	Inicio  int64
	Tamanio int64
	Nombre  [16]byte
}

// Inicializar Recibe un puntero Particion para ser modificado.
func (p *Particion) Inicializar(estado byte, tipo byte, fit byte, inicio int64,
	tamanio int64, nombre string) {
	p.Estado = estado
	p.Tipo = tipo
	p.Fit = fit
	p.Inicio = inicio
	p.Tamanio = tamanio
	copy(p.Nombre[:], nombre)
	for i := len(nombre); i < 16; i++ {
		p.Nombre[i] = byte(" "[0])
	}
}

// GetEstado recibe una copia de Particion ya que no necesita modificarlo.
func (p Particion) GetEstado() byte {
	return p.Estado
}

// SetEstado recibe un puntero Particion para ser modificado.
func (p *Particion) SetEstado(estado byte) {
	p.Estado = estado
}

// GetTipo retorna el valor de tipo
func (p Particion) GetTipo() byte {
	return p.Tipo
}

// SetTipo asigna el tipo
func (p *Particion) SetTipo(tipo byte) {
	p.Tipo = tipo
}

// GetFit retorna el fit de la particion
func (p Particion) GetFit() byte {
	return p.Fit
}

// SetFit asigna el fit
func (p *Particion) SetFit(fit byte) {
	p.Fit = fit
}

// GetInicio retorna la posicion inicial de la particion
func (p Particion) GetInicio() int64 {
	return p.Inicio
}

// SetInicio asigna la posicion inicial de la particion
func (p *Particion) SetInicio(inicio int64) {
	p.Inicio = inicio
}

// GetTamanio obtiene el tamaño de la particion
func (p Particion) GetTamanio() int64 {
	return p.Tamanio
}

// SetTamanio asigna el tamaño a la particion
func (p *Particion) SetTamanio(tamanio int64) {
	p.Tamanio = tamanio
}

// GetNombre retorna el nombre de la particion
func (p Particion) GetNombre() string {
	nombre := ""
	for i := 0; i < len(p.Nombre); i++ {
		if p.Nombre[i] != 0 {
			nombre += string(p.Nombre[i])
		}
	}
	return strings.TrimSpace(nombre)
}

// SetNombre asigna el nombre a la particion
func (p *Particion) SetNombre(nombre string) {
	copy(p.Nombre[:], nombre)
}
