package main

import (
	"fmt"
	"time"
)

func main() {
	// Ejemplo de inicialización
	usuario := Usuario{
		IDUsuario:         "U123",
		Nombre:            "Juan",
		Apellido:          "perez",
		CorreoElectronico: "juan.perez@example.com",
		Contrasena:        "securepassword",
		FechaRegistro:     time.Now(),
	}

	suscripcion := Suscripcion{
		IDSuscripcion:   "S456",
		IDUsuario:       usuario.IDUsuario,
		TipoPlan:        "Premium",
		FechaInicio:     time.Now(),
		FechaExpiracion: time.Now().AddDate(0, 1, 0), // Añadir un mes
	}

	perfil := Perfil{
		IDPerfil:               "P789",
		IDUsuario:              usuario.IDUsuario,
		NombrePerfil:           "Perfil Familiar",
		HistorialVisualizacion: "Serie A, Película B",
	}

	contenido := Contenido{
		IDContenido:  "C001",
		Titulo:       "Película Épica",
		Descripcion:  "Una aventura épica.",
		Genero:       "Aventura",
		FechaEstreno: time.Now().AddDate(-1, 0, 0),
		Duracion:     "2h30m",
		Buscador:     "épica",
	}

	recomendacion := Recomendaciones{
		IDRecomendacion:    "R123",
		IDPerfil:           perfil.IDPerfil,
		IDContenido:        contenido.IDContenido,
		FechaRecomendacion: time.Now(),
	}

	// Imprimir los datos
	fmt.Println("Usuario:", usuario.Nombre, usuario.Apellido)
	fmt.Println("Correo Electrónico:", usuario.CorreoElectronico)
	fmt.Println("Fecha de Registro:", usuario.FechaRegistro)
	fmt.Println("Suscripción:", suscripcion.TipoPlan)
	fmt.Println("Fecha de Expiración:", suscripcion.FechaExpiracion)
	fmt.Println("Perfil:", perfil.NombrePerfil)
	fmt.Println("Historial de Visualización:", perfil.HistorialVisualizacion)
	fmt.Println("Recomendación:", recomendacion.IDRecomendacion)
	fmt.Println("Contenido:", contenido.Titulo)
	fmt.Println("Género:", contenido.Genero)
	fmt.Println("Duración:", contenido.Duracion)

}
