package main

import (
	"fmt"
	"log"
	"os"

	"github.com/axolotl-go/axo-pages-server/controller"
	"github.com/axolotl-go/axo-pages-server/db"
	"github.com/axolotl-go/axo-pages-server/middleware"
	"github.com/axolotl-go/axo-pages-server/model"
	"github.com/axolotl-go/axo-pages-server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file :", err)
	}

	PORT := "3000"
	if p := os.Getenv("PORT"); p != "" {
		PORT = p
	}

	db.Dbconnection()
	db.DB.AutoMigrate(model.User{})

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	app.Static("/assets", "./public/assets", fiber.Static{
		Compress:  false,
		ByteRange: true,
		Browse:    false,
	})

	app.Static("/css", "./public/css")
	app.Static("/js", "./public/js")
	app.Static("/images", "./public/images")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html", false)
	})

	app.Post("/upload", controller.Uploader)

	app.Post("/register", routes.CreateUser)

	app.Post("/login", routes.LoginUser)

	app.Post("/logout", controller.LogOut)

	app.Get("/profile", middleware.Protected(), func(c *fiber.Ctx) error {
		user := c.Locals("user")
		return c.JSON(fiber.Map{"user": user})
	})

	app.Get("/:page", func(c *fiber.Ctx) error {
		page := c.Params("page")
		filepath := fmt.Sprintf("./public/%s.html", page)

		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			return controller.NotFound(c)
		}
		return c.SendFile(filepath, false)
	})

	app.Use("*", func(c *fiber.Ctx) error {
		return controller.NotFound(c)
	})

	app.Listen("0.0.0.0:" + PORT)
}
