package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func mountSuperAdminRoutes(app *fiber.App) {
	superAdmin := app.Group("/superadmin", func(c *fiber.Ctx) error {
		fmt.Println("SuperAdmin Middleware")
		return c.Next()
	})
	superAdmin.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("SuperAdmin Test")
	})

}
