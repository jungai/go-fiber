package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupMiddleware(a *fiber.App) {
	a.Use(
		// cors
		cors.New(),

		// rate limit
		limiter.New(limiter.Config{
			Max:        5,
			Expiration: 1 * time.Minute,
		}),

		// logger
		logger.New(),
	)
}
