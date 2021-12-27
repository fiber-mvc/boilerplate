package services

import (
	"github.com/fiber-mvc/fiber-mvc/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

type TestService struct {
	service.Service
}

func (s *TestService) Test() string {
	log.Println("TestService.Test")
	return "Hello World from TestService"
}

func (s *TestService) Setup(app *fiber.App) {
	log.Println("TestService setup")
}
