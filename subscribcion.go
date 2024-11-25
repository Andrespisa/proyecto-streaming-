package main

import "time"

// Suscripcion representa el plan de suscripción de un usuario.
type Suscripcion struct {
	IDSuscripcion   string
	IDUsuario       string
	TipoPlan        string
	FechaInicio     time.Time
	FechaExpiracion time.Time
}
