package main

import (
	"fiber-tina-store-API/configs"
	"fiber-tina-store-API/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()
	routes.UserRoute(app)

	app.Listen(":4040")
}
