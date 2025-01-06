package routes

import (
	"orderly/internal/api/middleware"
	"orderly/internal/domain/constants"
	"orderly/internal/domain/models"
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func mountAdminRoutes(app *fiber.App, handlers *di.Handlers) {
	admin := app.Group("/admin",
		middleware.ValidateJWT,
		middleware.ValidateAnyOfTheseRoles(constants.RoleAdmin),
	)
	{
		// account:= admin.Group("/account")
		// {
		// 	account.Post("/change-password", handlers.Handler.ChangePassword)
		// }

		users := admin.Group("/users", middleware.HavePrivilege(constants.Privilege_User_Management))
		{
			users.Get("", handlers.Handler.GetUsers)
			users.Get("/:id", handlers.Handler.GetUserByID)
			users.Patch("/activate/:id", handlers.Handler.ActivateByUUID(models.Users_TableName))
			users.Patch("/deactivate/:id", handlers.Handler.DeactivateByUUID(models.Users_TableName))
			users.Delete("/:id", handlers.Handler.SoftDeleteRecordByUUID(models.Users_TableName))
			users.Patch("/undo-delete/:id", handlers.Handler.UndoSoftDeleteRecordByUUID(models.Users_TableName))
		}

		category := admin.Group("/category", middleware.HavePrivilege(constants.Privilege_Inventory_Manager))
		{
			category.Post("", handlers.Handler.CreateCategory)
			category.Get("", handlers.Handler.GetCategories)
			category.Get("/:id", handlers.Handler.GetCategoryByID)
			category.Put("/:id", handlers.Handler.UpdateCategoryByID)
			category.Delete("/:id", handlers.Handler.SoftDeleteRecordByID(models.Category_TableName))
			category.Patch("/undo-delete/:id", handlers.Handler.UndoSoftDeleteRecordByID(models.Category_TableName))
		}

		product := admin.Group("/product", middleware.HavePrivilege(constants.Privilege_Inventory_Manager))
		{
			product.Post("", handlers.Handler.CreateProduct)
			product.Get("", handlers.Handler.GetProducts)
			product.Get("/:id", handlers.Handler.GetProductByID)
			product.Put("/:id", handlers.Handler.UpdateProductByID)
			product.Delete("/:id", handlers.Handler.SoftDeleteRecordByID(models.Products_TableName))
			product.Patch("/undo-delete/:id", handlers.Handler.UndoSoftDeleteRecordByID(models.Products_TableName))
			product.Get("/stock/:id", handlers.Handler.GetProductStockByID)
			product.Put("/stock/add/:id", handlers.Handler.AddProductStockByID)
		}

		order := admin.Group("/order", middleware.HavePrivilege(constants.Privilege_Sales_Manager))
		{
			order.Get("", handlers.Handler.GetOrders)
			order.Get("/:id", handlers.Handler.GetOrderDetails)
			order.Patch("/cancel/:id", handlers.Handler.CancelOrder)
			order.Patch("/mark-as-delivered/:id", handlers.Handler.MarkOrderAsDelivered)
		}

	}
}
