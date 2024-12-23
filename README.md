Proyecto: SPREEM - Sistema de Streaming
Descripción
Este proyecto es una aplicación web de un sistema de streaming llamado SPREEM. La aplicación permite a los usuarios registrarse, iniciar sesión, gestionar perfiles, suscripciones, y consumir contenidos. Está desarrollado en Go (Golang) con conexión a una base de datos MySQL.
Características principales
•	Autenticación de usuarios: Registro e inicio de sesión con verificación de credenciales.
•	Gestión de usuarios: Creación y administración de perfiles asociados a una cuenta.
•	Gestión de contenidos: Almacenamiento y despliegue de contenido multimedia con detalles como título, descripción, género, y fecha de estreno.
•	Suscripciones: Administración de planes de suscripción con opciones de pago.
•	Recomendaciones: Generación de recomendaciones personalizadas basadas en perfiles.
Configuración del Proyecto
1.	Clonar el repositorio:
2.	git clone <url-del-repositorio>
cd <nombre-del-directorio>
3.	Configurar la base de datos:
o	Asegúrate de tener un servidor MySQL en ejecución.
o	Actualiza las credenciales de la base de datos en el archivo conectar.go.
o	username := "root" // Usuario de MySQL
o	password := ""     // Contraseña de MySQL
o	host := "127.0.0.1"
o	port := "3306"
database := "StreamingDatabase"
o	Crea la base de datos con el nombre especificado (StreamingDatabase).
4.	Instalar dependencias: Ejecuta el siguiente comando:
go mod tidy
5.	Ejecutar la aplicación:
go run main.go
6.	Abrir la aplicación: Accede a la URL http://localhost:8080 desde tu navegador.

Estructura del Proyecto
El proyecto está organizado de la siguiente manera:
Archivos HTML (ubicados en templates/)
•	home.html: Página principal de la aplicación.
•	login.html: Página de inicio de sesión.
•	register.html: Página de registro de usuarios.
•	perfil.html: Gestión de perfiles de usuarios.
•	contenido.html: Visualización de contenidos disponibles.
•	planes.html: Descripción de los planes de suscripción.
•	subcripciones.html: Gestión de las suscripciones activas.
Archivos Go
•	main.go: Archivo principal que inicia el servidor web y gestiona las rutas.
•	usuario.go: Gestión de usuarios, incluyendo registro y autenticación.
•	conectar.go: Conexión y configuración de la base de datos MySQL.
•	contenido.go: Gestión de los contenidos multimedia.
•	perfil.go: Gestión de los perfiles asociados a usuarios.
•	suscripcion.go: Gestión de las suscripciones y planes de pago.
•	recomendaciones.go: Generación y administración de recomendaciones de contenido.
Base de Datos
La base de datos MySQL incluye las siguientes tablas:
Tabla Usuarios
CREATE TABLE IF NOT EXISTS `StreamingDatabase`.`Usuarios` (
    `IDUsuario` INT NOT NULL AUTO_INCREMENT,
    `Nombre` VARCHAR(100) NOT NULL,
    `Apellido` VARCHAR(100) NOT NULL,
    `Email` VARCHAR(100) UNIQUE NOT NULL,
    `Password` VARCHAR(255) NOT NULL,
    `FechaCreacion` DATETIME NOT NULL,
    PRIMARY KEY (`IDUsuario`)
) ENGINE = InnoDB;
Tabla Contenidos
CREATE TABLE IF NOT EXISTS `StreamingDatabase`.`Contenidos` (
    `IDContenido` INT NOT NULL AUTO_INCREMENT,
    `Titulo` VARCHAR(200) NOT NULL,
    `Descripcion` TEXT NOT NULL,
    `Genero` VARCHAR(50) NOT NULL,
    `Duracion` TIME NOT NULL,
    `Tags` VARCHAR(255),
    `FechaEstreno` DATETIME NOT NULL,
    PRIMARY KEY (`IDContenido`)
) ENGINE = InnoDB;
Tabla Perfiles
CREATE TABLE IF NOT EXISTS `StreamingDatabase`.`Perfiles` (
    `IDPerfil` INT NOT NULL AUTO_INCREMENT,
    `IDUsuario` INT NOT NULL,
    `Nombre` VARCHAR(100) NOT NULL,
    `Preferencias` TEXT,
    PRIMARY KEY (`IDPerfil`),
    FOREIGN KEY (`IDUsuario`) REFERENCES `StreamingDatabase`.`Usuarios`(`IDUsuario`) ON DELETE CASCADE
) ENGINE = InnoDB;
Tabla Recomendaciones
CREATE TABLE IF NOT EXISTS `StreamingDatabase`.`Recomendaciones` (
    `IDRecomendacion` INT NOT NULL AUTO_INCREMENT,
    `IDPerfil` INT NOT NULL,
    `IDContenido` INT NOT NULL,
    `FechaRecomendada` DATETIME NOT NULL,
    PRIMARY KEY (`IDRecomendacion`),
    FOREIGN KEY (`IDPerfil`) REFERENCES `StreamingDatabase`.`Perfiles`(`IDPerfil`) ON DELETE CASCADE,
    FOREIGN KEY (`IDContenido`) REFERENCES `StreamingDatabase`.`Contenidos`(`IDContenido`) ON DELETE CASCADE
) ENGINE = InnoDB;
Tabla Suscripciones
CREATE TABLE IF NOT EXISTS `StreamingDatabase`.`Suscripciones` (
    `IDSuscripcion` INT NOT NULL AUTO_INCREMENT,
    `IDUsuario` INT NOT NULL,
    `IDPlan` INT NOT NULL,
    `Monto` DECIMAL(10, 2) NOT NULL,
    `Tipo` ENUM('Mensual', 'Anual', 'Gratis') NOT NULL,
    `FechaInicio` DATETIME NOT NULL,
    `FechaExpiracion` DATETIME NOT NULL,
    PRIMARY KEY (`IDSuscripcion`),
    FOREIGN KEY (`IDUsuario`) REFERENCES `StreamingDatabase`.`Usuarios`(`IDUsuario`) ON DELETE CASCADE,
    FOREIGN KEY (`IDPlan`) REFERENCES `StreamingDatabase`.`PlanesPago`(`IDPlan`) ON DELETE CASCADE
) ENGINE = InnoDB;
Configuración
1.	Instalar Go en tu sistema.
2.	Configurar la conexión a MySQL en el archivo conectar.go.
3.	Crear las tablas en la base de datos usando los scripts proporcionados.
4.	Ejecutar el proyecto con:
go run main.go
Requisitos
•	Go 1.18 o superior.
•	MySQL Server.
Créditos
Desarrollado por: Andrés Sebastián Piedra Sánche
                  Josué Alexis Piedra Sánchez

