package routes

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type ApiRouter struct{}

func (ApiRouter) Setup(app *fiber.App) {
	log.Println("ApiRouter setup")

	group := app.Group("/api")

	group.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World! from api",
		})
	})
}
