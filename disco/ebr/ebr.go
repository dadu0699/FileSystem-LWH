package ebr

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
func (e *EBR) Inicializar(Estado byte, Fit byte, Inicio int64, Tamanio int64,
	Siguiente int64, Nombre string) {
	e.Estado = Estado
	e.Fit = Fit
	e.Inicio = Inicio
	e.Tamanio = Tamanio
	e.Siguiente = Siguiente
	copy(e.Nombre[:], Nombre)
}

// GetEstado recibe una copia de EBR ya que no necesita modificarlo.
func (e EBR) GetEstado() byte {
	return e.Estado
}

// SetEstado recibe un puntero EBR para ser modificado.
func (e *EBR) SetEstado(Estado byte) {
	e.Estado = Estado
}

// GetFit retorna el Fit del EBR
func (e EBR) GetFit() byte {
	return e.Fit
}

// SetFit asigna el Fit
func (e *EBR) SetFit(Fit byte) {
	e.Fit = Fit
}

// GetInicio retorna la posicion inicial del EBR
func (e EBR) GetInicio() int64 {
	return e.Inicio
}

// SetInicio asigna la posicion inicial del EBR
func (e *EBR) SetInicio(Inicio int64) {
	e.Inicio = Inicio
}

// GetTamanio obtiene el tamaño del EBR
func (e EBR) GetTamanio() int64 {
	return e.Tamanio
}

// SetTamanio asigna el tamaño al EBR
func (e *EBR) SetTamanio(Tamanio int64) {
	e.Tamanio = Tamanio
}

// GetSiguiente obtiene el Inicio del Siguiente EBR
func (e EBR) GetSiguiente() int64 {
	return e.Siguiente
}

// SetSiguiente asigna el Inicio del Siguiente EBR
func (e *EBR) SetSiguiente(Siguiente int64) {
	e.Siguiente = Siguiente
}

// GetNombre retorna el Nombre del EBR
func (e EBR) GetNombre() string {
	Nombre := ""
	for i := 0; i < len(e.Nombre); i++ {
		Nombre += string(e.Nombre[i])
	}
	return Nombre
}

// SetNombre asigna el Nombre al EBR
func (e *EBR) SetNombre(Nombre string) {
	copy(e.Nombre[:], Nombre)
}
