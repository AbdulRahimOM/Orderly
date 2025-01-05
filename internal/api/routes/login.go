package routes

import (
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func mountLoginRoutes(app *fiber.App, handlers *di.Handlers) {
	login := app.Group("/login")
	login.Post("/superAdmin", handlers.AccountHandler.SuperAdminSignin)
	login.Post("/admin", handlers.AccountHandler.AdminSignin)

	user := app.Group("/user")
	{
		user.Post("/signin", handlers.AccountHandler.UserSignIn)
		user.Post("/signup-get-otp", handlers.AccountHandler.UserSignUpGetOTP)
	}
}
