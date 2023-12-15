package utils

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// StartServer func for a starting a simple server
func StartServer(a *fiber.App) {
	//Build fiber connection url
	fiberConnURL, _ := ConnectionURLBuilder("fiber")

	//Run server
	err := a.Listen(fiberConnURL)
	if err != nil {
		log.Printf("Server is not running! Reason: '%v'", err)
	}
}
