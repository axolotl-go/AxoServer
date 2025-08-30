package http

import (
	"fmt"
	"os"

	"github.com/axolotl-go/axo-pages-server/controller"
	"github.com/axolotl-go/axo-pages-server/internal/auth"
	"github.com/axolotl-go/axo-pages-server/internal/user"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes monta todas las rutas en el app de Fiber
func SetupRoutes(app *fiber.App) {
	// Archivos estáticos
	app.Static("/assets", "./public/assets")
	app.Static("/css", "./public/css")
	app.Static("/js", "./public/js")
	app.Static("/images", "./public/images")

	// Rutas públicas
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html", false)
	})

	// User endpoints
	app.Post("/register", user.CreateUser)
	app.Post("/login", user.LoginUser)
	app.Post("/logout", user.LogOut)

	// Protected endpoints
	app.Get("/profile", auth.Protected(), func(c *fiber.Ctx) error {
		u := c.Locals("user")
		return c.JSON(fiber.Map{"user": u})
	})

	// Upload
	app.Post("/upload", controller.Uploader)

	// Rutas dinámicas de páginas
	app.Get("/:page", func(c *fiber.Ctx) error {
		page := c.Params("page")
		filepath := fmt.Sprintf("./public/%s.html", page)

		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			return NotFound(c)
		}
		return c.SendFile(filepath, false)
	})

	// Catch-all 404
	app.Use(func(c *fiber.Ctx) error {
		return NotFound(c)
	})
}
