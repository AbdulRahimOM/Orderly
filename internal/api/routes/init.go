package routes

import (
	"orderly/internal/infrastructure/db"
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	handlers := di.GetHandlers(db.DB)
	mountLoginRoutes(app, handlers)
	mountSuperAdminRoutes(app, handlers)
	mountAdminRoutes(app, handlers)
	mountUserRoutes(app, handlers)
	mountBrowseRoutes(app, handlers)
}
