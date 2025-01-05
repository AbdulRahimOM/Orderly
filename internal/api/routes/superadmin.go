package routes

import (
	"orderly/internal/api/middleware"
	"orderly/internal/domain/constants"
	"orderly/internal/domain/models"
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func mountSuperAdminRoutes(app *fiber.App, handlers *di.Handlers) {
	superAdmin := app.Group("/superAdmin",
		middleware.ValidateJWT,
		middleware.ValidateAnyOfTheseRoles(constants.RoleSuperAdmin),
	)

	admin := superAdmin.Group("/admin")
	{
		admin.Post("", handlers.Handler.CreateAdmin)
		admin.Get("", handlers.Handler.GetAdmins)
		admin.Get("/:id", handlers.Handler.GetAdminByID)
		admin.Put("/:id", handlers.Handler.UpdateAdminByID)
		admin.Delete("/:id", handlers.Handler.SoftDeleteRecordByUUID(models.Admins_TableName))
		admin.Patch("/undo-delete/:id", handlers.Handler.UndoSoftDeleteRecordByUUID(models.Admins_TableName))
		admin.Patch("/activate/:id", handlers.Handler.ActivateByUUID(models.Admins_TableName))
		admin.Patch("/deactivate/:id", handlers.Handler.DeactivateByUUID(models.Admins_TableName))
	}
}
