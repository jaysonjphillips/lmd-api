package storage

import "github.com/gofiber/fiber/v2"

func AddMovingBoxRoutes(router fiber.Router) {
	box := router.Group("/boxes")
	box.Post("/", CreateBox)
}
