package user

import (
	"github.com/axolotl-go/axo-pages-server/internal/auth"
	"github.com/axolotl-go/axo-pages-server/internal/db"
	"github.com/axolotl-go/axo-pages-server/internal/hash"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	hashedPassword, err := hash.Hash(user.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}
	user.Password = hashedPassword

	if err := db.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.JSON(
		fiber.Map{
			"message": "User created successfully",
		},
	)

}

func LoginUser(c *fiber.Ctx) error {
	var user User
	var login Login

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := db.DB.Where("username = ?", login.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	if !hash.CompareHash(user.Password, login.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	token, err := auth.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})

	return c.JSON(
		fiber.Map{
			"token": token,
		},
	)
}

func LogOut(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}
