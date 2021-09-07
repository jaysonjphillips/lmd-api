package wine

import "github.com/gofiber/fiber/v2"

func AddWineRoutes(router fiber.Router) {
	wineApi := router.Group("/wines")
	wineApi.Get("/", GetWines)
	wineApi.Get("/:id", GetWine)
	wineApi.Post("/", AddWine)
	wineApi.Put("/:id", UpdateWine)
	wineApi.Delete("/:id", DeleteWine)
}
