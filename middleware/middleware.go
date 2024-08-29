package middleware

import (
	"fiber-api/utils"

	"github.com/gofiber/fiber/v2"
)

// Middleware
func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// _, err := utils.VerifyToken(token)
	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	// ctx.Locals("userInfo", claims)
	// ctx.Locals("role", claims["role"])

	// Untuk melanjutkan ke fungsi berikutnya (handler)
	return ctx.Next()

	// if token != "secret" {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "unauthenticated",
	// 	})
	// }
	// return ctx.Next()
}

// Anothe middleware
func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}
