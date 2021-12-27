package controller

import (
	"github.com/fiber-mvc/fiber-mvc/controller"
	"github.com/gofiber/fiber/v2"
	"log"
)

type TestController struct {
	controller.Controller
}

func (c *TestController) Setup(app *fiber.App) {
	log.Println("TestController Setup")
}
