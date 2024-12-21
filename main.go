package main

import (
	"fmt"
	"net/http"
	"time"
)

// Función principal
func main() {
	// Inicializar datos de ejemplo
	initSampleData()

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

// Sirve la página principal
func serveHomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
	<!DOCTYPE html>
	<html lang="es">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>SPREEM - Inicio</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 0;
				padding: 0;
				background-color:rgb(0, 0, 0);
				color: #333;
			}
			header {
				background-color:rgb(129, 0, 0);
				color: white;
				padding: 10px 20px;
				text-align: center;
			}
			header nav a {
				color: white;
				text-decoration: none;
				margin: 0 15px;
				font-size: 1.2em;
			}
			header nav a:hover {
				text-decoration: underline;
			}
			.container {
				text-align: center;
				padding: 50px 20px;
			}
			h1 {
				font-size: 3em;
				margin-bottom: 10px;
				color: #444;
			}
			p {
				font-size: 1.2em;
				margin-bottom: 20px;
			}
			.button {
				display: inline-block;
				margin: 10px;
				padding: 15px 30px;
				font-size: 1em;
				color: #fff;
				background-color:rgb(129, 0, 0);
				border: none;
				border-radius: 5px;
				cursor: pointer;
				text-decoration: none;
			}
			.button:hover {
				background-color:rgb(129, 0, 0);
			}
		</style>
	</head>
	<body>
		<header>
			<h1>SPREEM</h1>
			<nav>
				<a href="/">Inicio</a>
				<a href="/contenidos">Contenido</a>
				<a href="/perfiles">perfiles</a>
			</nav>
		</header>
		<div class="container">
			<h1>Bienvenido a SPREEM</h1>
			<p>Tu plataforma para películas y series</p>
			<a href="/login" class="button">Iniciar Sesión</a>
			<a href="/register" class="button">Regístrate</a>
		</div>
	</body>
	</html>
	`
	w.Write([]byte(html))
}
func handleContenidos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
	<!DOCTYPE html>
	<html lang="es">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Contenidos</title>
		<style>
			body { font-family: Arial, sans-serif; background-color:rgb(0, 0, 0); color: #333; margin: 0; padding: 0; }
			header { background-color:rgb(129, 0, 0); color: white; padding: 20px; text-align: center; }
			h1 { margin: 0; font-size: 2.5em; }
			.catalogo { display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 20px; padding: 20px; }
			.card { background: white; padding: 15px; border-radius: 5px; box-shadow: 0 4px 6px rgba(129, 0, 0, 0.1); }
			.card h2 { font-size: 1.4em; margin-bottom: 10px; color:rgb(129, 0, 0); }
			.card p { margin: 5px 0; color: #666; }
		</style>
	</head>
	<body>
		<header>
			<h1>Catálogo de Contenidos</h1>
		</header>
		<div class="catalogo">
	`
	for _, contenido := range contenidos {
		html += fmt.Sprintf(`
			<div class="card">
				<h2>%s</h2>
				<p><strong>Género:</strong> %s</p>
				<p><strong>Duración:</strong> %s</p>
				<p><strong>Descripción:</strong> %s</p>
			</div>
		`, contenido.Titulo, contenido.Genero, contenido.Duracion, contenido.Descripcion)
	}

	html += `
		</div>
	</body>
	</html>
	`
	w.Write([]byte(html))
}
func handlePerfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
	<!DOCTYPE html>
	<html lang="es">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Perfiles</title>
		<style>
			body { font-family: Arial, sans-serif; background-color:rgb(0, 0, 0); color: #333; margin: 0; padding: 0; }
			header { background-color:rgb(129, 0, 0); color: white; padding: 20px; text-align: center; }
			h1 { margin: 0; font-size: 2.5em; }
			.container { padding: 20px; }
			ul { list-style: none; padding: 0; }
			li { background: white; margin: 10px 0; padding: 10px; border-radius: 5px; box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); }
		</style>
	</head>
	<body>
		<header>
			<h1>Lista de Perfiles</h1>
		</header>
		<div class="container">
			<ul>
	`
	for _, perfil := range perfiles {
		html += fmt.Sprintf(`
			<li>
				<strong>ID:</strong> %s <br>
				<strong>Nombre:</strong> %s <br>
				<strong>Usuario ID:</strong> %s <br>
				<strong>Preferencias:</strong> %s
			</li>
		`, perfil.IDPerfil, perfil.Nombre, perfil.IDUsuario, perfil.Preferencias)
	}

	html += `
			</ul>
		</div>
	</body>
	</html>
	`
	w.Write([]byte(html))
}

// Sirve la página de inicio de sesión
func serveLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
	<!DOCTYPE html>
	<html lang="es">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Iniciar Sesión - SPREEM</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color:rgb(0, 0, 0);
				color: #333;
				text-align: center;
				padding: 50px 20px;
			}
			h1 {
				font-size: 2.5em;
				margin-bottom: 20px;
			}
			form {
				display: inline-block;
				text-align: left;
			}
			input {
				display: block;
				margin: 10px 0;
				padding: 10px;
				width: 100%;
				font-size: 1em;
			}
			button {
				padding: 10px 20px;
				background-color:rgb(129, 0, 0);
				color: white;
				border: none;
				border-radius: 5px;
				cursor: pointer;
				font-size: 1em;
			}
			button:hover {
				background-color:rgb(129, 0, 0);
			}
		</style>
	</head>
	<body>
		<h1>Iniciar Sesión</h1>
		<p>Accede a tu cuenta para disfrutar de SPREEM.</p>
		<form action="/login" method="post">
			<input type="text" name="username" placeholder="Usuario" required>
			<input type="password" name="password" placeholder="Contraseña" required>
			<button type="submit">Acceder</button>
		</form>
	</body>
	</html>
	`
	w.Write([]byte(html))
}

// Sirve la página de registro
func serveRegisterPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
	<!DOCTYPE html>
	<html lang="es">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Regístrate - SPREEM</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color:rgb(0, 0, 0);
				color: #333;
				text-align: center;
				padding: 50px 20px;
			}
			h1 {
				font-size: 2.5em;
				margin-bottom: 20px;
			}
			form {
				display: inline-block;
				text-align: left;
			}
			input {
				display: block;
				margin: 10px 0;
				padding: 10px;
				width: 100%;
				font-size: 1em;
			}
			button {
				padding: 10px 20px;
				background-color:rgb(129, 0, 0);
				color: white;
				border: none;
				border-radius: 5px;
				cursor: pointer;
				font-size: 1em;
			}
			button:hover {
				background-color:rgb(129, 0, 0);
			}
		</style>
	</head>
	<body>
		<h1>Regístrate</h1>
		<p>Crea tu cuenta y comienza a disfrutar de SPREEM.</p>
		<form action="/register" method="post">
			<input type="text" name="username" placeholder="Usuario" required>
			<input type="email" name="email" placeholder="Correo Electrónico" required>
			<input type="password" name="password" placeholder="Contraseña" required>
			<button type="submit">Registrarse</button>
		</form>
	</body>
	</html>
	`
	w.Write([]byte(html))
}

// Inicializar datos de ejemplo
func initSampleData() {
	// Crear un usuario de ejemplo
	usuario := NewUsuario("U123", "Juan", "Pérez", "juan.perez@example.com", "securepassword", time.Now())
	usuarios = append(usuarios, *usuario)

	// Crear contenido de ejemplo
	contenido := NewContenido("C001", "Película Épica", "Una aventura épica.", "Aventura", "2h30m", "épica", time.Now().AddDate(-1, 0, 0))
	contenidos = append(contenidos, *contenido)

	// Crear un perfil de ejemplo
	perfil := CrearPerfil("P789", usuario.IDUsuario, "Perfil Familiar", "Serie A, Película B")
	perfiles = append(perfiles, perfil)

	// Crear una recomendación de ejemplo
	recomendacion := CrearRecomendacion("R123", perfil.IDPerfil, contenido.IDContenido, time.Now())
	recomendaciones = append(recomendaciones, recomendacion)
}
