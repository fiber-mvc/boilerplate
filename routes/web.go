package routes

import (
	"boilerplate/app/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"log"
)

type WebRouter struct{}

func (WebRouter) Setup(app *fiber.App) {

	log.Println("WebRouter setup")

	homeController := di.GetInstance("HomeController").(*controllers.HomeController)

	app.Get("/", homeController.Index)
	app.Post("/", homeController.Post)
}
