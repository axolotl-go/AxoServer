myapi/
├── cmd/
│ └── server/
│ └── main.go # Punto de entrada
├── internal/
│ ├── auth/
│ │ ├── jwt.go # Generación y validación de JWT
│ │ └── middleware.go # Middleware de autenticación
│ ├── hash/
│ │ └── bcrypt.go # Hash y verificación de contraseñas
│ ├── user/
│ │ ├── model.go # Definición de struct User
│ │ ├── repository.go # Acceso a BD para Users
│ │ └── service.go # Registro, login
│ ├── http/
│ │ ├── router.go # Definición de rutas
│ │ ├── handlers_user.go # Endpoints de user
│ │ └── handlers_host.go # Endpoints de host
│ └── database/
│ └── postgres.go # Conexión a Postgres (puede ser otro)
├── go.mod
└── go.sum

//
myapi/
├── cmd/
│ └── server/
│ └── main.go
├── config/
│ └── config.go
├── internal/
│ ├── auth/
│ ├── hash/
│ ├── host/
│ ├── user/
│ ├── http/
│ └── database/
├── public/ # Archivos estáticos
│ ├── index.html
│ ├── styles.css
│ └── js/
│ └── app.js
├── go.mod
└── go.sum
