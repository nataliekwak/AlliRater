package main

import (
	"go-admin/database"
	"go-admin/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()

	// Browser blocks requests from different ports
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:4200",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
