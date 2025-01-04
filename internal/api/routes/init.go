package routes

import (
	"orderly/internal/infrastructure/db"
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	handlers := di.GetHandlers(db.PublicDB)
	mountLoginRoutes(app, handlers)
	mountSuperAdminRoutes(app, handlers)
}
