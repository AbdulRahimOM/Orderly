package routes

import (
	"orderly/internal/api/middleware"
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func mountSuperAdminRoutes(app *fiber.App, handlers *di.Handlers) {
	superAdmin := app.Group("/superAdmin",
		middleware.ValidateJWT,
		// middleware.ValidateRole("superadmin"),
	)

	admin := superAdmin.Group("/admin")
	{
		admin.Post("", handlers.AccountHandler.CreateAdmin)
	}

}
