package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
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

// Función para registrar un usuario en la base de datos con contraseña encriptada
func registerUsuario(nombre, username, apellido, email, password string) error {
	// Encriptar la contraseña antes de almacenarla
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err // Si ocurre un error en la encriptación, devolverlo
	}

	query := "INSERT INTO usuarios (nombre, username, apellido, email, password, FechaCreacion) VALUES (?, ?, ?, ?, ?, ?)"
	fechaCreacion := time.Now() // Genera la fecha y hora actual

	// Ejecutar la consulta en la base de datos
	_, err = getDB().Exec(query, nombre, username, apellido, email, string(hashedPassword), fechaCreacion)
	return err
}

// Datos iniciales de usuarios
var usuarios []Usuario

// Función para verificar las credenciales de inicio de sesión
func loginUsuario(email, password string) (*Usuario, error) {
	// Obtener el usuario con el email proporcionado
	query := "SELECT id_usuario, nombre, username, apellido, email, password, FechaCreacion FROM usuarios WHERE email = ?"
	row := getDB().QueryRow(query, email)

	var usuario Usuario
	// Escanear el resultado de la consulta
	err := row.Scan(&usuario.IDUsuario, &usuario.Nombre, &usuario.Username, &usuario.Apellido, &usuario.Email, &usuario.Password, &usuario.FechaCreacion)
	if err != nil {
		if err == sql.ErrNoRows {
			// Si no se encuentra el usuario
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, err // Otro error de consulta
	}

	// Comparar la contraseña proporcionada con la almacenada en la base de datos
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// Si las contraseñas no coinciden
			return nil, fmt.Errorf("credenciales incorrectas")
		}
		return nil, err // Otro error de comparación
	}

	// Si la contraseña es correcta, devolver el usuario
	return &usuario, nil
}
