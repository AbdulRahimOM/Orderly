package main

import (
	"fmt"
	"log"
	route "orderly/internal/api/routes"
	"orderly/internal/infrastructure/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Configs.Env.CORSAllowedOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowCredentials: true, // Enable credentials for specific origins
	}))

	// health check
	app.Get("/health", healthCheck)

	// mount other routes
	route.MountRoutes(app)

	err := app.Listen(fmt.Sprintf(":%s", config.Configs.Env.Port))
	if err != nil {
		log.Fatal("Couldn't start the server. Error: " + err.Error())
	}
}
