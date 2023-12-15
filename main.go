package main

import (
	"book-restapi/pkg/configs"
	"book-restapi/pkg/middleware"
	"book-restapi/pkg/routes"
	"book-restapi/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//Define fiber config
	config := configs.FiberConfig()

	//Define a new fiber app with config
	app := fiber.New(config)

	//Middlewares
	middleware.FiberMiddleware(app)

	//Routes
	routes.PublicRoutes(app)

	//Start server
	utils.StartServer(app)
}
