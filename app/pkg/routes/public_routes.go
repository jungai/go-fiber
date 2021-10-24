package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jungai/first-fiber-go/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/", controllers.HelloWorld)
	route.Get("/error", controllers.GetError)
}
