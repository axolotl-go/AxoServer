package controller

import (
	"fmt"
	"os"
	"path/filepath"

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

func Uploader(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).SendString("Error al recibir archivo")
	}

	uploadPath := filepath.Join("public", file.Filename)
	if err := c.SaveFile(file, uploadPath); err != nil {
		return c.Status(500).SendString("Error al guardar archivo")
	}

	return c.JSON(fiber.Map{
		"message": "Archivo subido correctamente",
		"url":     fmt.Sprintf("/%s", file.Filename),
	})

}

func LogOut(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}
