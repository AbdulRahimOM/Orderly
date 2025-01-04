package routes

import (
	"orderly/internal/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func mountSuperAdminRoutes(app *fiber.App) {
	superAdmin := app.Group("/superadmin",
	middleware.ValidateJWT)
	superAdmin.Post("/login", )
}
