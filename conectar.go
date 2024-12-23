package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // Importar el driver MySQL
)

// Variable global para la conexión
var db *sql.DB

// Conectar a la base de datos MySQL
func conect() error {
	username := "root" // Cambia por tu usuario de MySQL
	password := ""     // Agrega tu contraseña aquí si es necesaria
	host := "127.0.0.1"
	port := "3306"
	database := "StreamingDatabase"

	// Crear el string de conexión
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	var err error

	// Abrir la conexión a la base de datos
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error al abrir la conexión: %v", err)
	}

	// Verificar la conexión
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error al conectar con la base de datos: %v", err)
	}

	// Configurar el tiempo de espera para las conexiones inactivas
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10) // Máximo de conexiones abiertas simultáneamente
	db.SetMaxIdleConns(5)  // Máximo de conexiones inactivas

	fmt.Println("Conexión exitosa a la base de datos MySQL")

	// Crear la tabla `usuarios` si no existe
	if err := verificarCrearTablaUsuarios(); err != nil {
		return fmt.Errorf("error al verificar/crear la tabla 'usuarios': %v", err)
	}

	return nil
}

// Cerrar la conexión a la base de datos
func cerrar() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Println("Error al cerrar la conexión de la base de datos:", err)
		} else {
			fmt.Println("Conexión cerrada correctamente.")
		}
	}
}

// Obtener la conexión de la base de datos
func getDB() *sql.DB {
	return db
}

// Verificar y crear la tabla `usuarios` si no existe
func verificarCrearTablaUsuarios() error {
	query := `
    CREATE TABLE IF NOT EXISTS usuarios (
        id_usuario INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(255) NOT NULL,
        username VARCHAR(255) NOT NULL,
        apellido VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        FechaCreacion DATETIME NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    `
	// Intentar crear la tabla si no existe
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error al crear/verificar la tabla 'usuarios': %v", err)
	}

	fmt.Println("Tabla 'usuarios' y la columna 'FechaCreacion' verificados correctamente.")
	return nil
}

// verificar y crea la tabla de suscripciones' si no existe
func verificarCrearTablaSuscripciones() error {
	query := `
    CREATE TABLE IF NOT EXISTS suscripciones (
        id_suscripcion INT AUTO_INCREMENT PRIMARY KEY,
        id_usuario INT NOT NULL,
        tipo_plan VARCHAR(255) NOT NULL,
        fecha_inicio DATETIME NOT NULL,
        fecha_expiracion DATETIME NOT NULL,
        FOREIGN KEY (id_usuario) REFERENCES usuarios(id_usuario)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    `
	// Intentar crear la tabla si no existe
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error al crear/verificar la tabla 'suscripciones': %v", err)
	}

	fmt.Println("Tabla 'suscripciones' verificada correctamente.")
	return nil
}

// Verificar y crear la tabla `contenidos` si no existe
func verificarCrearTablaContenidos() error {
	query := `
    CREATE TABLE IF NOT EXISTS contenidos (
        idcontenido INT AUTO_INCREMENT PRIMARY KEY,
        titulo VARCHAR(255) NOT NULL,
        descripcion TEXT NOT NULL,
        genero VARCHAR(100),
        duracion VARCHAR(50),
        tags VARCHAR(255),
        fechaestreno DATETIME NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    `
	// Intentar crear la tabla si no existe
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error al crear/verificar la tabla 'contenidos': %v", err)
	}

	fmt.Println("Tabla 'contenidos' verificada correctamente.")
	return nil
}
