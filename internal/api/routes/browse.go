package routes

import (
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func mountBrowseRoutes(app *fiber.App, handlers *di.Handlers) {
	browse := app.Group("/browse")
	{
		browse.Get("/category", handlers.Handler.GetCategories)
		browse.Get("/category/:id", handlers.Handler.GetCategoryByID)
		browse.Get("/product", handlers.Handler.GetProducts)
		browse.Get("/product/:id", handlers.Handler.GetProductByID)
	}
}
