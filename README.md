# MyAPI – Plataforma para Hosting de Páginas Web

MyAPI es una aplicación escrita en **Go** que permite alojar y gestionar páginas web de manera sencilla, inspirada en plataformas como **Vercel**, **Cloudflare Pages** y **GitHub Pages**.
Está diseñada con una arquitectura modular y escalable, ideal para proyectos modernos de despliegue de aplicaciones estáticas y dinámicas.

---

## Estructura del Proyecto

```
myapi/
├── cmd/
│   └── server/
│       └── main.go        # Punto de entrada principal de la aplicación
├── config/
│   └── config.go          # Configuración general de la app (variables, entorno)
├── internal/              # Lógica interna y módulos principales
│   ├── auth/              # Manejo de autenticación y sesiones
│   ├── hash/              # Funciones de hashing (contraseñas, tokens)
│   ├── host/              # Lógica para alojar y servir páginas web
│   ├── user/              # Gestión de usuarios (registro, login, perfiles)
│   ├── http/              # Handlers HTTP y ruteo con Fiber
│   └── database/          # Conexión y acceso a base de datos con GORM
├── public/                # Archivos estáticos servidos al cliente
│   ├── index.html
│   ├── styles.css
│   └── js/
│       └── app.js
├── go.mod                 # Dependencias del proyecto
└── go.sum                 # Checksum de dependencias
```

---

## Instalación y Ejecución

1. Clona el repositorio:

   ```bash
   git clone https://github.com/axolotl-go/AxoServer.git
   cd AxoServer
   ```

2. Instala las dependencias:

   ```bash
   go mod tidy
   ```

3. Configura las variables de entorno (ejemplo en `.env`):

   ```env
   PORT=8080
   DB_DSN=mydatabase
   JWT_KEY=mysecretpassword
   ```

4. Ejecuta el servidor:

   ```bash
   go run ./cmd/server/main.go
   ```

5. Abre en el navegador:

   ```
   http://localhost:8080
   ```

---

## Funcionalidades

* Autenticación segura con JWT
* Gestión de usuarios (registro, login, perfiles)
* Hosting de páginas web estáticas en la carpeta `/public`
* Arquitectura modular fácil de extender
* Soporte para bases de datos relacionales (ej. SQLite, PostgreSQL, MySQL)
* Uso de **Fiber** como framework web rápido y eficiente
* Manejo de la base de datos con **GORM**

---

## Tecnologías

* Lenguaje: Go
* Framework web: Fiber
* ORM: GORM
* Base de datos: SQLite (extensible a MySQL/PostgreSQL)
* Frontend estático: HTML, CSS, JS (servido desde `/public`)
