package user

import "github.com/gofiber/fiber/v2"

func AddUserRoutes(router fiber.Router) {
	userApi := router.Group("/users")
	userApi.Get("/:id", GetUser)
	userApi.Post("/", CreateUser)
	userApi.Put("/:id", UpdateUser)
	userApi.Delete("/:id", DeleteUser)
}
