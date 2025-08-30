package main

import (
	"log"
	"os"

	"github.com/axolotl-go/axo-pages-server/internal/db"
	"github.com/axolotl-go/axo-pages-server/internal/http"
	"github.com/axolotl-go/axo-pages-server/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var port = "3000"

func init() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
}

func main() {
	err := db.DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	http.SetupRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
