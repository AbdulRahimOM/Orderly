package routes

import (
	"orderly/internal/api/middleware"
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func mountLoginRoutes(app *fiber.App, handlers *di.Handlers) {
	login := app.Group("/login")
	{
		login.Post("/superAdmin", handlers.Handler.SuperAdminSignin)
		login.Post("/admin", handlers.Handler.AdminSignin)
		login.Post("/user", handlers.Handler.UserSignIn)
	}

	app.Post("/user-signup-get-otp", handlers.Handler.UserSignUpGetOTP)
	app.Post("/user-signup-verify-otp", middleware.ValidateJWT, handlers.Handler.UserSignUpVerifyOTP)
}
