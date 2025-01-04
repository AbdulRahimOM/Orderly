package routes

import (
	"orderly/internal/api/middleware"
	"orderly/internal/domain/models"
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
		admin.Get("", handlers.AccountHandler.GetAdmins)
		admin.Get("/:id", handlers.AccountHandler.GetAdminByID)
		admin.Put("/:id", handlers.AccountHandler.UpdateAdminByID)
		admin.Delete("/:id", handlers.AccountHandler.SoftDeleteRecordByID(models.Admins_TableName))
		admin.Patch("/undo-delete/:id", handlers.AccountHandler.UndoSoftDeleteRecordByID(models.Admins_TableName))
		admin.Patch("/block/:id", handlers.AccountHandler.BlockByID(models.Admins_TableName))
		admin.Patch("/unblock/:id", handlers.AccountHandler.UnblockByID(models.Admins_TableName))
	}
}
