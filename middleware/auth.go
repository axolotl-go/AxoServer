package middleware

import (
	"os"
	"strings"

	"github.com/axolotl-go/axo-pages-server/utils"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var token string

		// 1️⃣ Revisar en el header
		authHeader := c.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// 2️⃣ Si no hay en header, revisar cookie
		if token == "" {
			token = c.Cookies("token")
		}

		// 3️⃣ Si sigue vacío → error
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// 4️⃣ Validar JWT
		claims, err := utils.ParseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Guardamos info del usuario en c.Locals()
		c.Locals("user", claims)

		return c.Next()
	}
}
