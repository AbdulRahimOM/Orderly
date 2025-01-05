package routes

import (
	"orderly/internal/infrastructure/db"

	"github.com/gofiber/fiber/v2"
)

func MountDevRoutes(app *fiber.App){
	app.Get("/reset-db",db.ResetDB)
}
