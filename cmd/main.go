package main

import (
	"fmt"
	"log"
	"orderly/internal/api/controls"
	route "orderly/internal/api/routes"
	"orderly/internal/infrastructure/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func healthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"msg": "ok",
	})
}

func main() {

	app := fiber.New(fiber.Config{
		AppName:       "Ordely",
		StrictRouting: true,
	})
	app.Use(logger.New())

	// health check
	app.Get("/health", healthCheck)

	// mount other routes
	route.MountRoutes(app)

	// schedule rate change cron job
	controls.ScheduleMidNightRateChangeOperation()

	err := app.Listen(fmt.Sprintf(":%s", config.Configs.Env.Port))
	if err != nil {
		log.Fatal("Couldn't start the server. Error: " + err.Error())
	}
}
