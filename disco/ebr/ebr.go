package ebr

// EBR modelo de la estructura
type EBR struct {
	estado    byte
	fit       byte
	inicio    int64
	tamanio   int64
	siguiente int64
	nombre    [16]byte
}

// Inicializar Recibe un puntero EBR para ser modificado.
func (e *EBR) Inicializar(estado byte, fit byte, inicio int64, tamanio int64,
	siguiente int64, nombre string) {
	e.estado = estado
	e.fit = fit
	e.inicio = inicio
	e.tamanio = tamanio
	e.siguiente = siguiente
	copy(e.nombre[:], nombre)
}

// GetEstado recibe una copia de EBR ya que no necesita modificarlo.
func (e EBR) GetEstado() byte {
	return e.estado
}

// SetEstado recibe un puntero EBR para ser modificado.
func (e *EBR) SetEstado(estado byte) {
	e.estado = estado
}

// GetFit retorna el Fit del EBR
func (e EBR) GetFit() byte {
	return e.fit
}

// SetFit asigna el fit
func (e *EBR) SetFit(fit byte) {
	e.fit = fit
}

// GetInicio retorna la posicion inicial del EBR
func (e EBR) GetInicio() int64 {
	return e.inicio
}

// SetInicio asigna la posicion inicial del EBR
func (e *EBR) SetInicio(inicio int64) {
	e.inicio = inicio
}

// GetTamanio obtiene el tamaño del EBR
func (e EBR) GetTamanio() int64 {
	return e.tamanio
}

// SetTamanio asigna el tamaño al EBR
func (e *EBR) SetTamanio(tamanio int64) {
	e.tamanio = tamanio
}

// GetSiguiente obtiene el inicio del siguiente EBR
func (e EBR) GetSiguiente() int64 {
	return e.siguiente
}

// SetSiguiente asigna el inicio del siguiente EBR
func (e *EBR) SetSiguiente(siguiente int64) {
	e.siguiente = siguiente
}

// GetNombre retorna el nombre del EBR
func (e EBR) GetNombre() string {
	nombre := ""
	for i := 0; i < len(e.nombre); i++ {
		nombre += string(e.nombre[i])
	}
	return nombre
}

// SetNombre asigna el nombre al EBR
func (e *EBR) SetNombre(nombre string) {
	copy(e.nombre[:], nombre)
}
