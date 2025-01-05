package routes

import (
	"orderly/internal/domain/models"
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func mountAdminRoutes(app *fiber.App, handlers *di.Handlers) {
	admin := app.Group("/admin")
	{
		// account:= admin.Group("/account")
		// {
		// 	account.Post("/change-password", handlers.Handler.ChangePassword)
		// }

		users := admin.Group("/users")
		{
			users.Get("", handlers.Handler.GetUsers)
			users.Get("/:id", handlers.Handler.GetUserByID)
			users.Patch("/activate/:id", handlers.Handler.ActivateByUUID(models.Users_TableName))
			users.Patch("/deactivate/:id", handlers.Handler.DeactivateByUUID(models.Users_TableName))
			users.Delete("/:id", handlers.Handler.SoftDeleteRecordByUUID(models.Users_TableName))
			users.Patch("/undo-delete/:id", handlers.Handler.UndoSoftDeleteRecordByUUID(models.Users_TableName))
		}

	}
}
