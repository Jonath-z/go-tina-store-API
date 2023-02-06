package routes

import (
	"fiber-tina-store-API/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(app *fiber.App){
   app.Post("/product", controllers.CreateProduct)
}