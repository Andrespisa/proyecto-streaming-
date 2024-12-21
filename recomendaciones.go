package main

import (
	"time"
)

// Definición de la estructura Recomendacion
type Recomendacion struct {
	IDRecomendacion  string
	IDPerfil         string
	IDContenido      string
	FechaRecomendada time.Time
}

// Constructor para Recomendacion
func CrearRecomendacion(idRecomendacion, idPerfil, idContenido string, fechaRecomendada time.Time) Recomendacion {
	return Recomendacion{idRecomendacion, idPerfil, idContenido, fechaRecomendada}
}

// Variable global para almacenar recomendaciones
var recomendaciones []Recomendacion
