package main

import "time"

// Usuario representa un usuario del sistema.
type Usuario struct {
	IDUsuario         string
	Nombre            string
	Apellido          string
	CorreoElectronico string
	Contrasena        string
	FechaRegistro     time.Time
}
