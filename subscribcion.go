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

// NewSuscripcion es un constructor para la estructura Suscripcion.
func NewSuscripcion(tipoPlan string) *Suscripcion {
	return &Suscripcion{
		TipoPlan: tipoPlan,
	}
}

// GetTipoPlan obtiene el tipo de plan de la suscripción.
func (s *Suscripcion) GetTipoPlan() string {
	return s.TipoPlan
}

// SetTipoPlan establece el tipo de plan de la suscripción.
func (s *Suscripcion) SetTipoPlan(tipoPlan string) {
	s.TipoPlan = tipoPlan
}

// EsPremium verifica si la suscripción es Premium.
func (s *Suscripcion) EsPremium() bool {
	return s.TipoPlan == "Premium"
}
