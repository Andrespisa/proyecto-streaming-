package main

// Definición de la estructura Perfil
type Perfil struct {
	IDPerfil     string
	IDUsuario    string
	Nombre       string
	Preferencias string
}

// Constructor para Perfil
func CrearPerfil(idPerfil, idUsuario, nombre, preferencias string) Perfil {
	return Perfil{idPerfil, idUsuario, nombre, preferencias}
}

// Datos iniciales de perfiles
var perfiles []Perfil
