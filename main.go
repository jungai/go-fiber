package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jungai/first-fiber-go/app/pkg/middlewares"
	"github.com/jungai/first-fiber-go/app/pkg/routes"
)

func main() {
	app := fiber.New()

	// register middleware
	middlewares.SetupMiddleware(app)

	// register routes
	routes.PublicRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
