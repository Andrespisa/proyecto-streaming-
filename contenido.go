package main

import "time"

// Contenido representa el contenido disponible en la plataforma.
type Contenido struct {
	IDContenido  string
	Titulo       string
	Descripcion  string
	Genero       string
	FechaEstreno time.Time
	Duracion     string
	Buscador     string
}
