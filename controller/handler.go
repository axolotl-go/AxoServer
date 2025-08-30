package controller

import (
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

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
