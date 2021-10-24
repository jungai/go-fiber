package controllers

import "github.com/gofiber/fiber/v2"

func GetError(c *fiber.Ctx) error {
	return fiber.NewError(404, "Custom Error NotFound Kub")
}
