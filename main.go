package main

import (
	"fmt"
	"os"
	"phillz-life-dashboard/auth"
	"phillz-life-dashboard/storage"
	"phillz-life-dashboard/user"
	"phillz-life-dashboard/wine"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public")

	// setup default api version/base
	api := app.Group("/api").Group("/v1")

	//add wine routes
	wine.AddWineRoutes(api)
	user.AddUserRoutes(api)
	auth.AddAuthRoutes(api)
	storage.AddMovingBoxRoutes(api)

	port := os.Getenv("PORT")

	//add sneaker routes
	app.Listen(fmt.Sprintf(":%s", port))
}
