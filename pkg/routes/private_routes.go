package routes

import (
	"book-restapi/app/controllers"
	"book-restapi/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App, w *controllers.WW) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/book", middleware.JWTProtected(), w.CreateBook)
}
