package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/loveyandex/go-mongo-celestia-scan/controller"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// Default config
	app.Use(cors.New())

	//routes adn controllers
	controller.NewUserController(app)

	//finish to listen
	app.Listen(":9595")
}
