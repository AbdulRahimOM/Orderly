package routes

import (
	"orderly/internal/api/middleware"
	"orderly/internal/domain/constants"
	"orderly/internal/domain/models"
	"orderly/internal/infrastructure/di"

	"github.com/gofiber/fiber/v2"
)

func mountUserRoutes(app *fiber.App, handlers *di.Handlers) {
	user := app.Group("/user",
		middleware.ValidateJWT,
		middleware.ValidateAnyOfTheseRoles(constants.RoleUser),
	)
	{
		account := user.Group("/account")
		{

			profile := account.Group("/profile")
			{
				profile.Get("", handlers.Handler.GetUserProfile)
			}

			address := account.Group("/address")
			{
				address.Get("", handlers.Handler.GetUserAddresses)
				address.Get("/:id", handlers.Handler.GetUserAddressByID)
				address.Post("", handlers.Handler.CreateUserAddress)
				address.Put("/:id", handlers.Handler.UpdateUserAddressByID)
				address.Delete("/:id", handlers.Handler.HardDeleteRecordByUUID(models.Addresses_TableName))
			}
		}

		cart := user.Group("/cart")
		{
			cart.Get("", handlers.Handler.GetCart)
			cart.Put("", handlers.Handler.AddToCart)
			cart.Delete("/product/:id", handlers.Handler.RemoveProductFromCart)
			cart.Patch("/update-quantity", handlers.Handler.UpdateCartItemQuantity)
			cart.Delete("/clear", handlers.Handler.ClearCart)
		}

	}

}
