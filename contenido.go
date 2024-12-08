package main

import (
	"errors"
	"time"
)

// Contenido representa el contenido disponible en la plataforma.
type Contenido struct {
	idContenido  string
	titulo       string
	descripcion  string
	genero       string
	fechaEstreno time.Time
	duracion     string
	buscador     string
}

// NewContenido es un constructor que crea un nuevo contenido
func NewContenido(id, titulo, descripcion, genero, duracion, buscador string, fechaEstreno time.Time) *Contenido {
	return &Contenido{
		idContenido:  id,
		titulo:       titulo,
		descripcion:  descripcion,
		genero:       genero,
		fechaEstreno: fechaEstreno,
		duracion:     duracion,
		buscador:     buscador,
	}
}

// Métodos Getter (accesores)
func (c *Contenido) GetIDContenido() string {
	return c.idContenido
}

func (c *Contenido) GetTitulo() string {
	return c.titulo
}

func (c *Contenido) GetDescripcion() string {
	return c.descripcion
}

func (c *Contenido) GetGenero() string {
	return c.genero
}

func (c *Contenido) GetFechaEstreno() time.Time {
	return c.fechaEstreno
}

func (c *Contenido) GetDuracion() string {
	return c.duracion
}

func (c *Contenido) GetBuscador() string {
	return c.buscador
}

// Métodos Setter (mutadores)
func (c *Contenido) SetTitulo(titulo string) {
	c.titulo = titulo
}

func (c *Contenido) SetDescripcion(descripcion string) {
	c.descripcion = descripcion
}

func (c *Contenido) SetGenero(genero string) {
	c.genero = genero
}

func (c *Contenido) SetDuracion(duracion string) {
	c.duracion = duracion
}

func (c *Contenido) SetBuscador(buscador string) {
	c.buscador = buscador
}

// Validar si la fecha de estreno es correcta
func (c *Contenido) EsFechaValida() error {
	if c.fechaEstreno.After(time.Now()) {
		return errors.New("la fecha de estreno no puede ser futura")
	}
	return nil
}
