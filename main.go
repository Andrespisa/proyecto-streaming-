package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// Intentar conectar con la base de datos
	if err := conect(); err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer cerrar() // Asegúrate de cerrar la conexión al final

	// Configurar rutas
	http.HandleFunc("/", serveHomePage)
	http.HandleFunc("/login", serveLoginPage)
	http.HandleFunc("/register", serveRegisterPage)
	http.HandleFunc("/usuarios", handleUsuarios)
	http.HandleFunc("/contenidos", handleContenidos)
	http.HandleFunc("/perfiles", handlePerfiles)

	// Iniciar el servidor
	fmt.Println("Servidor web ejecutándose en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

// Cargar el template HTML desde la carpeta "templates"
func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	tmplPath := filepath.Join("templates", templateName) // Ruta a la plantilla
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error al cargar la página", http.StatusInternalServerError)
		return
	}

	// Ejecutar la plantilla con los datos proporcionados
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error al procesar la plantilla", http.StatusInternalServerError)
	}
}

// Función para la página de inicio
func serveHomePage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "SPREEM - Inicio",
	}
	renderTemplate(w, "home.html", data)
}

// Función para la página de login
func serveLoginPage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Iniciar Sesión - SPREEM",
	}
	renderTemplate(w, "login.html", data)
}

// Función para la página de registro
func serveRegisterPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Si el método es GET, muestra la página de registro
		data := struct {
			Title string
		}{
			Title: "Regístrate - SPREEM",
		}
		renderTemplate(w, "register.html", data)
	} else if r.Method == http.MethodPost {
		// Si el método es POST, procesa el registro del usuario
		nombre := r.FormValue("nombre")
		username := r.FormValue("username")
		apellido := r.FormValue("apellido")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if nombre == "" || username == "" || apellido == "" || email == "" || password == "" {
			http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
			return
		}

		// Registrar el usuario en la base de datos
		if err := registerUsuario(nombre, username, apellido, email, password); err != nil {
			http.Error(w, "Error al registrar el usuario: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Mostrar mensaje de éxito
		http.Redirect(w, r, "/contenidos", http.StatusSeeOther)

	} else {
		// Método no soportado
		http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)

	}
}

// Función para la página de perfil
func handlePerfiles(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Perfiles - SPREEM",
	}
	renderTemplate(w, "perfil.html", data)
}

///

// Función para manejar la página de Contenidos (ahora con template)
func handleContenidos(w http.ResponseWriter, r *http.Request) {
	// Obtenemos los contenidos de la base de datos
	contenidos, err := getContenidos(db)
	if err != nil {
		log.Printf("Error al consultar contenidos: %v", err)
		http.Error(w, fmt.Sprintf("Error al consultar contenidos: %v", err), http.StatusInternalServerError)
		return
	}

	// Estructura para pasar los datos al template
	data := struct {
		Title      string
		Contenidos []Contenido
	}{
		Title:      "Contenidos - SPREEM",
		Contenidos: contenidos,
	}

	// Renderizamos el template con los datos
	renderTemplate(w, "contenido.html", data)
}
