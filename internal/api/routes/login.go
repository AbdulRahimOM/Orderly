package routes

import (
	"orderly/internal/api/middleware"
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func mountLoginRoutes(app *fiber.App, handlers *di.Handlers) {
	login := app.Group("/login")
	{
		login.Post("/superAdmin", handlers.AccountHandler.SuperAdminSignin)
		login.Post("/admin", handlers.AccountHandler.AdminSignin)
		login.Post("/user", handlers.AccountHandler.UserSignIn)
	}

	app.Post("/user-signup-get-otp", handlers.AccountHandler.UserSignUpGetOTP)
	app.Post("/user-signup-verify-otp", middleware.ValidateJWT, handlers.AccountHandler.UserSignUpVerifyOTP)
}
