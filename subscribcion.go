package main

import (
	"fmt"
	"time"
)

// Suscripcion representa el plan de suscripción de un usuario.
type Suscripcion struct {
	IDSuscripcion   string    `json:"id_suscripcion"`
	IDUsuario       string    `json:"id_usuario"`
	TipoPlan        string    `json:"tipo_plan"`
	FechaInicio     time.Time `json:"fecha_inicio"`
	FechaExpiracion time.Time `json:"fecha_expiracion"`
}

// Función para registrar una suscripción en la base de datos
func registerSuscripcion(idUsuario, tipoPlan string, fechaInicio, fechaExpiracion time.Time) error {
	query := `
		INSERT INTO suscripciones (id_usuario, tipo_plan, fecha_inicio, fecha_expiracion)
		VALUES (?, ?, ?, ?)
	`

	_, err := getDB().Exec(query, idUsuario, tipoPlan, fechaInicio, fechaExpiracion)
	if err != nil {
		return fmt.Errorf("error al registrar la suscripción: %v", err)
	}
	return nil
}

// Función para obtener las suscripciones de un usuario
func getSuscripcionesDeUsuario(idUsuario string) ([]Suscripcion, error) {
	query := `SELECT id_suscripcion, id_usuario, tipo_plan, fecha_inicio, fecha_expiracion
			  FROM suscripciones WHERE id_usuario = ?`

	rows, err := getDB().Query(query, idUsuario)
	if err != nil {
		return nil, fmt.Errorf("error al obtener las suscripciones: %v", err)
	}
	defer rows.Close()

	var suscripciones []Suscripcion
	for rows.Next() {
		var s Suscripcion
		if err := rows.Scan(&s.IDSuscripcion, &s.IDUsuario, &s.TipoPlan, &s.FechaInicio, &s.FechaExpiracion); err != nil {
			return nil, fmt.Errorf("error al leer los resultados de la consulta: %v", err)
		}
		suscripciones = append(suscripciones, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante el recorrido de filas: %v", err)
	}

	return suscripciones, nil
}
