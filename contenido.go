package main

import (
	"time"
)

// Definición de la estructura Contenido
type Contenido struct {
	IDContenido  string
	Titulo       string
	Descripcion  string
	Genero       string
	Duracion     string
	Tags         string
	FechaEstreno time.Time
}

// Constructor para Contenido
func NewContenido(id, titulo, descripcion, genero, duracion, tags string, fechaEstreno time.Time) *Contenido {
	return &Contenido{id, titulo, descripcion, genero, duracion, tags, fechaEstreno}
}

// Datos iniciales de contenidos
var contenidos []Contenido
