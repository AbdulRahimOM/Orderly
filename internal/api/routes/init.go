package routes

import (
	"orderly/internal/infrastructure/db"
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	handlers:=di.GetHandlers(db.PublicDB)
	mountLoginRoutes(app,handlers)
	// for _, val := range app.Stack() {
	// 	for _, route := range val{
	// 		fmt.Println(route)
	// 	}
	// }
}
