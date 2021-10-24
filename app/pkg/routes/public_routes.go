package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jungai/first-fiber-go/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	// v1
	route := a.Group("/api/v1")
	route.Get("/", controllers.HelloWorld)
	route.Get("/questions", controllers.GetQuestions)
	route.Get("/error", controllers.GetError)
}
