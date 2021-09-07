package auth

import "github.com/gofiber/fiber/v2"

func AddAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/login", Login)
}
