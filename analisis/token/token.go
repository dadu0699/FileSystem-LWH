package token

// Token modelo de la estructura
type Token struct {
	idToken int
	fila    int
	columna int
	tipo    string
	valor   string
}

// Inicializar  Recibe un puntero Token para ser modificado.
func (t *Token) Inicializar(idToken int, fila int, columna int,
	tipo string, valor string) {
	t.idToken = idToken
	t.fila = fila
	t.columna = columna
	t.tipo = tipo
	t.valor = valor
}

// SetID recibe un puntero Token para ser modificado.
func (t *Token) SetID(idToken int) {
	t.idToken = idToken
}

// GetID recibe una copia de Token ya que no necesita modificarlo.
func (t Token) GetID() int {
	return t.idToken
}

// SetFila asigna el numero de fila
func (t *Token) SetFila(fila int) {
	t.fila = fila
}

// GetFila retorna el numero de fila
func (t Token) GetFila() int {
	return t.fila
}

// SetColumna asigna el numero de columna
func (t *Token) SetColumna(columna int) {
	t.columna = columna
}

// GetColumna retorna el numero de columna
func (t Token) GetColumna() int {
	return t.columna
}

// SetTipo asigna el tipo
func (t *Token) SetTipo(tipo string) {
	t.tipo = tipo
}

// GetTipo retorna el tipo
func (t Token) GetTipo() string {
	return t.tipo
}

// SetValor asigna el valor
func (t *Token) SetValor(valor string) {
	t.valor = valor
}

// GetValor retorna el valor
func (t Token) GetValor() string {
	return t.valor
}
