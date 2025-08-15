package main

import (
	"fmt"
	"os"

	"github.com/axolotl-go/axo-pages-server/controller"
	"github.com/axolotl-go/axo-pages-server/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	PORT := "3000"
	if p := os.Getenv("PORT"); p != "" {
		PORT = p
	}

	db.Dbconnection()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Static("/assets", "./public/assets", fiber.Static{
		Compress:  false,
		ByteRange: true,
		Browse:    false,
	})

	app.Static("/css", "./public/css")
	app.Static("/js", "./public/js")
	app.Static("/images", "./public/images")

	app.Post("/upload", controller.Uploader)

	app.Get("/:page", func(c *fiber.Ctx) error {
		page := c.Params("page")
		filepath := fmt.Sprintf("./public/%s.html", page)

		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			return controller.NotFound(c)
		}
		return c.SendFile(filepath, false)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html", false)
	})

	app.Use("*", func(c *fiber.Ctx) error {
		return controller.NotFound(c)
	})

	app.Listen(":" + PORT)
}
