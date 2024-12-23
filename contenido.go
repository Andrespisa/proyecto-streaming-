package main

import (
	"database/sql"
	"fmt"
	"time"
)

// Estructura para los contenidos
type Contenido struct {
	IDContenido  int       `json:"idcontenido"`
	Titulo       string    `json:"titulo"`
	Descripcion  string    `json:"descripcion"`
	Genero       string    `json:"genero"`
	Duracion     string    `json:"duracion"`
	Tags         string    `json:"tags"`
	FechaEstreno time.Time `json:"fechaEstreno"`
}

// Obtener todos los contenidos de la base de datos
func getContenidos(db *sql.DB) ([]Contenido, error) {
	query := "SELECT idcontenido, titulo, descripcion, genero, duracion, tags, fechaestreno FROM contenidos"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SQL: %v", err)
	}
	defer rows.Close()

	var contenidos []Contenido
	for rows.Next() {
		var contenido Contenido
		var fechaEstreno []byte // Usar []byte para manejar la fecha antes de convertirla
		err := rows.Scan(&contenido.IDContenido, &contenido.Titulo, &contenido.Descripcion, &contenido.Genero, &contenido.Duracion, &contenido.Tags, &fechaEstreno)
		if err != nil {
			return nil, fmt.Errorf("error al mapear los resultados: %v", err)
		}

		// Convertir []byte a time.Time
		if len(fechaEstreno) > 0 {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", string(fechaEstreno))
			if err != nil {
				return nil, fmt.Errorf("error al convertir fecha: %v", err)
			}
			contenido.FechaEstreno = parsedTime
		}

		// Imprimir para depurar
		fmt.Printf("Contenido: %+v\n", contenido)
		contenidos = append(contenidos, contenido)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante la iteración de resultados: %v", err)
	}

	return contenidos, nil
}
