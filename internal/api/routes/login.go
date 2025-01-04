package routes

import (
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func mountLoginRoutes(app *fiber.App, handlers *di.Handlers) {
	login := app.Group("/login")
	login.Post("/superAdmin", handlers.AccountHandler.SuperAdminSignin)
}
