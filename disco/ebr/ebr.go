package ebr

import "strings"

// EBR modelo de la estructura
type EBR struct {
	Estado    byte
	Fit       byte
	Inicio    int64
	Tamanio   int64
	Siguiente int64
	Nombre    [16]byte
}

// Inicializar Recibe un puntero EBR para ser modificado.
func (e *EBR) Inicializar(estado byte, fit byte, inicio int64, tamanio int64,
	siguiente int64, nombre string) {
	e.Estado = estado
	e.Fit = fit
	e.Inicio = inicio
	e.Tamanio = tamanio
	e.Siguiente = siguiente
	copy(e.Nombre[:], nombre)
	for i := len(nombre); i < 16; i++ {
		e.Nombre[i] = byte(" "[0])
	}
}

// GetEstado recibe una copia de EBR ya que no necesita modificarlo.
func (e EBR) GetEstado() byte {
	return e.Estado
}

// SetEstado recibe un puntero EBR para ser modificado.
func (e *EBR) SetEstado(estado byte) {
	e.Estado = estado
}

// GetFit retorna el Fit del EBR
func (e EBR) GetFit() byte {
	return e.Fit
}

// SetFit asigna el Fit
func (e *EBR) SetFit(fit byte) {
	e.Fit = fit
}

// GetInicio retorna la posicion inicial del EBR
func (e EBR) GetInicio() int64 {
	return e.Inicio
}

// SetInicio asigna la posicion inicial del EBR
func (e *EBR) SetInicio(inicio int64) {
	e.Inicio = inicio
}

// GetTamanio obtiene el tamaño del EBR
func (e EBR) GetTamanio() int64 {
	return e.Tamanio
}

// SetTamanio asigna el tamaño al EBR
func (e *EBR) SetTamanio(tamanio int64) {
	e.Tamanio = tamanio
}

// GetSiguiente obtiene el Inicio del Siguiente EBR
func (e EBR) GetSiguiente() int64 {
	return e.Siguiente
}

// SetSiguiente asigna el Inicio del Siguiente EBR
func (e *EBR) SetSiguiente(siguiente int64) {
	e.Siguiente = siguiente
}

// GetNombre retorna el Nombre del EBR
func (e EBR) GetNombre() string {
	var nombre strings.Builder
	for i := 0; i < len(e.Nombre); i++ {
		if e.Nombre[i] != 0 {
			nombre.WriteString(string(e.Nombre[i]))
		}
	}
	return strings.TrimSpace(nombre.String())
}

// SetNombre asigna el Nombre al EBR
func (e *EBR) SetNombre(nombre string) {
	copy(e.Nombre[:], nombre)
}
