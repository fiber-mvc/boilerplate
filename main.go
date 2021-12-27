package main

import (
	"boilerplate/app/controllers"
	"boilerplate/app/models"
	"boilerplate/app/services"
	"boilerplate/routes"
	core "github.com/fiber-mvc/fiber-mvc"
	"github.com/fiber-mvc/fiber-mvc/database"
	"github.com/fiber-mvc/fiber-mvc/routing"

	"github.com/gofiber/fiber/v2"
)

func main() {

	core.Boot(core.Config{

		Database: database.DBConfig{
			Driver:   "sqlite",
			Database: "db.db",
			Models: []interface{}{
				&models.User{},
			},
		},

		Fiber: fiber.Config{},

		Services: map[string]interface{}{
			"TestService": &services.TestService{},
		},

		Controllers: map[string]interface{}{
			"HomeController": &controllers.HomeController{},
		},

		Routers: []routing.Router{
			&routes.WebRouter{},
			&routes.ApiRouter{},
		},
	})
}
