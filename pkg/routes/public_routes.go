package routes

import (
	"book-restapi/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/user/signup", controllers.UserSignUp)
}
