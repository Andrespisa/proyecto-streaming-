package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Definición de la estructura Usuario
type Usuario struct {
	IDUsuario     string
	Nombre        string
	Apellido      string
	Email         string
	Password      string
	FechaCreacion time.Time
}

// Constructor para Usuario
func NewUsuario(id, nombre, apellido, email, password string, fechaCreacion time.Time) *Usuario {
	return &Usuario{id, nombre, apellido, email, password, fechaCreacion}
}

// Manejo de usuarios (JSON)
func handleUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// Datos iniciales de usuarios
var usuarios []Usuario
