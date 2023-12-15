package routes

import (
	"book-restapi/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/books", controllers.GetBooks)
	route.Get("/book/:id", controllers.GetBook)

	route.Post("/user/signup", controllers.UserSignUp)
	route.Post("/user/signin", controllers.UserSignIn)
}
