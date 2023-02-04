package responses

import "github.com/gofiber/fiber/v2"

type UserResponse struct {
	Status  int        `JSON:"status"`
	Message string     `JSON:"string"`
	Data    *fiber.Map `JSON:"data"`
}
