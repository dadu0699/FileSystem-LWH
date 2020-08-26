package particion

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
func (p *Particion) Inicializar(Estado byte, Tipo byte, Fit byte, Inicio int64,
	Tamanio int64, Nombre string) {
	p.Estado = Estado
	p.Tipo = Tipo
	p.Fit = Fit
	p.Inicio = Inicio
	p.Tamanio = Tamanio
	copy(p.Nombre[:], Nombre)
}

// GetEstado recibe una copia de Particion ya que no necesita modificarlo.
func (p Particion) GetEstado() byte {
	return p.Estado
}

// SetEstado recibe un puntero Particion para ser modificado.
func (p *Particion) SetEstado(Estado byte) {
	p.Estado = Estado
}

// GetTipo retorna el valor de Tipo
func (p Particion) GetTipo() byte {
	return p.Tipo
}

// SetTipo asigna el Tipo
func (p *Particion) SetTipo(Tipo byte) {
	p.Tipo = Tipo
}

// GetFit retorna el Fit de la particion
func (p Particion) GetFit() byte {
	return p.Fit
}

// SetFit asigna el Fit
func (p *Particion) SetFit(Fit byte) {
	p.Fit = Fit
}

// GetInicio retorna la posicion inicial de la particion
func (p Particion) GetInicio() int64 {
	return p.Inicio
}

// SetInicio asigna la posicion inicial de la particion
func (p *Particion) SetInicio(Inicio int64) {
	p.Inicio = Inicio
}

// GetTamanio obtiene el tamaño de la particion
func (p Particion) GetTamanio() int64 {
	return p.Tamanio
}

// SetTamanio asigna el tamaño a la particion
func (p *Particion) SetTamanio(Tamanio int64) {
	p.Tamanio = Tamanio
}

// GetNombre retorna el Nombre de la particion
func (p Particion) GetNombre() string {
	Nombre := ""
	for i := 0; i < len(p.Nombre); i++ {
		Nombre += string(p.Nombre[i])
	}
	return Nombre
}

// SetNombre asigna el Nombre a la particion
func (p *Particion) SetNombre(Nombre string) {
	copy(p.Nombre[:], Nombre)
}
