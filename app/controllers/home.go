package controllers

import (
	"boilerplate/app/models"
	"boilerplate/app/services"
	"github.com/fiber-mvc/fiber-mvc/controller"
	"github.com/fiber-mvc/fiber-mvc/database"
	"github.com/fiber-mvc/fiber-mvc/validation"

	"github.com/gofiber/fiber/v2"
	"log"
)

type HomeController struct {
	controller.Controller
	testService *services.TestService `di.inject:"TestService"`
}

func (c *HomeController) Index(ctx *fiber.Ctx) error {
	log.Println("test:", c.testService.Test())
	return ctx.Render("index", map[string]interface{}{
		"Name": "World!",
	})
}

func (c *HomeController) Post(ctx *fiber.Ctx) error {
	payload := struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=3"`
	}{}

	if err := ctx.BodyParser(&payload); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid payload")
	}

	errors := validation.ValidateStruct(payload)
	if len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	user := models.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
	}
	database.DB.Create(&user)
	return ctx.JSON(user)
}

func (c *HomeController) Setup(app *fiber.App) {
	log.Println("HomeController setup")
}
