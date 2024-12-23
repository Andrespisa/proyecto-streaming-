package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Definición de la estructura Usuario
type Usuario struct {
	IDUsuario     string    `json:"id_usuario"`
	Nombre        string    `json:"nombre"`
	Username      string    `json:"username"`
	Apellido      string    `json:"apellido"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	FechaCreacion time.Time `json:"fecha_creacion"`
}

// Constructor para Usuario
func NewUsuario(id, nombre, username, apellido, email, password string, fechaCreacion time.Time) *Usuario {
	return &Usuario{
		IDUsuario:     id,
		Nombre:        nombre,
		Username:      username,
		Apellido:      apellido,
		Email:         email,
		Password:      password,
		FechaCreacion: fechaCreacion,
	}
}

// Manejo de usuarios (JSON)
func handleUsuarios(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Obtener lista de usuarios y devolverla como JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(usuarios)
		return
	}
	http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
}

// Función para registrar un usuario en la base de datos
func registerUsuario(nombre, username, apellido, email, password string) error {
	query := `INSERT INTO usuarios (nombre, username, apellido, email, password, FechaCreacion)
			  VALUES (?, ?, ?, ?, ?, ?)`
	fechaCreacion := time.Now() // Genera la fecha y hora actual

	// Ejecutar la consulta en la base de datos
	_, err := getDB().Exec(query, nombre, username, apellido, email, password, fechaCreacion)
	return err
}

// Datos iniciales de usuarios
var usuarios []Usuario
