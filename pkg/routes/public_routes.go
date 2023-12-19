package routes

import (
	"book-restapi/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App, w *controllers.WW) {
	route := a.Group("/api/v1")

	route.Get("/books", w.GetBooks)
	route.Get("/book/:id", w.GetBook)

	route.Post("/user/signup", controllers.UserSignUp)
	route.Post("/user/signin", controllers.UserSignIn)
}
