package routes

import (
	"github.com/axolotl-go/axo-pages-server/db"
	"github.com/axolotl-go/axo-pages-server/model"
	"github.com/axolotl-go/axo-pages-server/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	hashedPassword, err := utils.Hash(user.Password)

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
	var user model.User
	var login model.Login

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

	if !utils.CompareHash(user.Password, login.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	token, err := utils.GenerateJWT(user.ID, user.Username)
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
