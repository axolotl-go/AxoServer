package http

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	notFoundPath := "./public/NotFound.html"
	if _, err := os.Stat(notFoundPath); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{
			"error": "Page not found",
			"code":  404,
		})
	}

	return c.Status(404).SendFile(notFoundPath)
}
