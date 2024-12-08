package main

import (
	"fmt"
	"time"
)

func main() {
	// Crear un usuario utilizando el constructor
	usuario := NewUsuario("U123", "Juan", "Pérez", "juan.perez@example.com", "securepassword", time.Now())

	// Crear una suscripción asociada al usuario
	suscripcion := Suscripcion{
		IDSuscripcion:   "S456",
		IDUsuario:       usuario.GetIDUsuario(),
		TipoPlan:        "Premium",
		FechaInicio:     time.Now(),
		FechaExpiracion: time.Now().AddDate(0, 1, 0), // Añadir un mes
	}

	// Crear un perfil asociado al usuario
	perfil := Perfil{
		IDPerfil:               "P789",
		IDUsuario:              usuario.GetIDUsuario(),
		NombrePerfil:           "Perfil Familiar",
		HistorialVisualizacion: "Serie A, Película B",
	}

	// Crear contenido con información relevante
	contenido := NewContenido("C001", "Película Épica", "Una aventura épica.", "Aventura", "2h30m", "épica", time.Now().AddDate(-1, 0, 0))

	// Verificar si la fecha de estreno es válida
	if err := contenido.EsFechaValida(); err != nil {
		// Si la fecha de estreno es inválida, mostramos un mensaje de error
		fmt.Println("Error en la fecha de estreno:", err)
		return
	}

	// Crear recomendación para el perfil y contenido
	recomendacion := Recomendaciones{
		IDRecomendacion:    "R123",
		IDPerfil:           perfil.IDPerfil,
		IDContenido:        contenido.GetIDContenido(),
		FechaRecomendacion: time.Now(),
	}

	// Imprimir los datos del usuario
	fmt.Println("Usuario:", usuario.GetNombre(), usuario.GetApellido())
	fmt.Println("Correo Electrónico:", usuario.GetCorreoElectronico())
	fmt.Println("Fecha de Registro:", usuario.GetFechaRegistro())

	// Imprimir los datos de la suscripción
	fmt.Println("Suscripción:", suscripcion.TipoPlan)
	fmt.Println("Fecha de Expiración:", suscripcion.FechaExpiracion)

	// Imprimir los datos del perfil
	fmt.Println("Perfil:", perfil.NombrePerfil)
	fmt.Println("Historial de Visualización:", perfil.HistorialVisualizacion)

	// Imprimir la recomendación
	fmt.Println("Recomendación ID:", recomendacion.IDRecomendacion)

	// Imprimir los datos del contenido
	fmt.Println("Contenido:", contenido.GetTitulo())
	fmt.Println("Género:", contenido.GetGenero())
	fmt.Println("Duración:", contenido.GetDuracion())
}
