package main

import (
	"book-restapi/app/controllers"
	"book-restapi/app/queries"
	"book-restapi/pkg/configs"
	"book-restapi/pkg/routes"
	"book-restapi/pkg/utils"
	"book-restapi/platform/database"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	//Define fiber config
	config := configs.FiberConfig()

	//Define a new fiber app with config
	app := fiber.New(config)

	q, _ := database.PostgresSQLConnection()

	bq := queries.NewBookQueries(q)
	c := controllers.NewBookController(&database.Queries{
		UserQueries: nil,
		BookQueries: bq,
	})

	//Middlewares
	//middleware.FiberMiddleware(app)

	//Routes
	routes.PublicRoutes(app, c)
	routes.PrivateRoutes(app, c)
	routes.NotFoundRoute(app)

	//Start server
	utils.StartServer(app)
}
