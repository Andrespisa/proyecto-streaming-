package main

import "time"

// Usuario representa un usuario del sistema.
type Usuario struct {
	idUsuario         string
	nombre            string
	apellido          string
	correoElectronico string
	contrasena        string
	fechaRegistro     time.Time
}

// NewUsuario es un constructor que crea un nuevo usuario
func NewUsuario(id, nombre, apellido, correo, contrasena string, fechaRegistro time.Time) *Usuario {
	return &Usuario{
		idUsuario:         id,
		nombre:            nombre,
		apellido:          apellido,
		correoElectronico: correo,
		contrasena:        contrasena,
		fechaRegistro:     fechaRegistro,
	}
}

// Métodos Getter (accesores)
func (u *Usuario) GetIDUsuario() string {
	return u.idUsuario
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) GetApellido() string {
	return u.apellido
}

func (u *Usuario) GetCorreoElectronico() string {
	return u.correoElectronico
}

func (u *Usuario) GetContrasena() string {
	return u.contrasena
}

func (u *Usuario) GetFechaRegistro() time.Time {
	return u.fechaRegistro
}
