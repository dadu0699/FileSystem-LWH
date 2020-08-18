package errort

// ErrorT modelo de la estructura
type ErrorT struct {
	idError     int
	fila        int
	columna     int
	tipo        string
	caracter    string
	descripcion string
}

// Inicializar  Recibe un puntero ErrorT para ser modificado.
func (e *ErrorT) Inicializar(idError int, fila int, columna int,
	tipo string, caracter string, descripcion string) {
	e.idError = idError
	e.fila = fila
	e.columna = columna
	e.tipo = tipo
	e.caracter = caracter
	e.descripcion = descripcion
}

// SetID recibe un puntero ErrorT para ser modificado.
func (e *ErrorT) SetID(idError int) {
	e.idError = idError
}

// GetID recibe una copia de ErrorT ya que no necesita modificarlo.
func (e ErrorT) GetID() int {
	return e.idError
}

// SetFila asigna el numero de fila
func (e *ErrorT) SetFila(fila int) {
	e.fila = fila
}

// GetFila retorna el numero de fila
func (e ErrorT) GetFila() int {
	return e.fila
}

// SetColumna asigna el numero de columna
func (e *ErrorT) SetColumna(columna int) {
	e.columna = columna
}

// GetColumna retorna el numero de columna
func (e ErrorT) GetColumna() int {
	return e.columna
}

// SetTipo asigna el tipo
func (e *ErrorT) SetTipo(tipo string) {
	e.tipo = tipo
}

// GetTipo retorna el tipo
func (e ErrorT) GetTipo() string {
	return e.tipo
}

// SetCaracter asigna el valor
func (e *ErrorT) SetCaracter(caracter string) {
	e.caracter = caracter
}

// GetCaracter retorna el valor
func (e ErrorT) GetCaracter() string {
	return e.caracter
}

// SetDescripcion asigna el valor
func (e *ErrorT) SetDescripcion(descripcion string) {
	e.descripcion = descripcion
}

// GetDescripcion retorna el valor
func (e ErrorT) GetDescripcion() string {
	return e.descripcion
}
