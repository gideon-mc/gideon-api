package gideon

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Gideon API",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "localhost:5000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		Expiration:   1 * time.Minute,
		CacheControl: true,
	}))

	return app
}
