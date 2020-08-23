package particion

// Particion modelo de la estructura
type Particion struct {
	estado  byte
	tipo    byte
	fit     byte
	inicio  int64
	tamanio int64
	nombre  [16]byte
}

// Inicializar Recibe un puntero Particion para ser modificado.
func (p *Particion) Inicializar(estado byte, tipo byte, fit byte, inicio int64,
	tamanio int64, nombre string) {
	p.estado = estado
	p.tipo = tipo
	p.fit = fit
	p.inicio = inicio
	p.tamanio = tamanio
	copy(p.nombre[:], nombre)
}

// GetEstado recibe una copia de Particion ya que no necesita modificarlo.
func (p Particion) GetEstado() byte {
	return p.estado
}

// SetEstado recibe un puntero Particion para ser modificado.
func (p *Particion) SetEstado(estado byte) {
	p.estado = estado
}

// GetTipo retorna el valor de tipo
func (p Particion) GetTipo() byte {
	return p.tipo
}

// SetTipo asigna el tipo
func (p *Particion) SetTipo(tipo byte) {
	p.tipo = tipo
}

// GetFit retorna el Fit de la particion
func (p Particion) GetFit() byte {
	return p.fit
}

// SetFit asigna el fit
func (p *Particion) SetFit(fit byte) {
	p.fit = fit
}

// GetInicio retorna la posicion inicial de la particion
func (p Particion) GetInicio() int64 {
	return p.inicio
}

// SetInicio asigna la posicion inicial de la particion
func (p *Particion) SetInicio(inicio int64) {
	p.inicio = inicio
}

// GetTamanio obtiene el tamaño de la particion
func (p Particion) GetTamanio() int64 {
	return p.tamanio
}

// SetTamanio asigna el tamaño a la particion
func (p *Particion) SetTamanio(tamanio int64) {
	p.tamanio = tamanio
}

// GetNombre retorna el nombre de la particion
func (p Particion) GetNombre() string {
	nombre := ""
	for i := 0; i < len(p.nombre); i++ {
		nombre += string(p.nombre[i])
	}
	return nombre
}

// SetNombre asigna el nombre a la particion
func (p *Particion) SetNombre(nombre string) {
	copy(p.nombre[:], nombre)
}
