package di

import (
	accounthandler "orderly/internal/api/handler"
	repo "orderly/internal/repository"
	accountuc "orderly/internal/usecase/account"

	"gorm.io/gorm"
)

//dependency injection

type Handlers struct {
	AccountHandler *accounthandler.Handler
}

func GetHandlers(db *gorm.DB) *Handlers {
	repo := repo.NewPublicRepository(db)

	accountUsecase := accountuc.NewAccountUsecase(repo)

	accountHandler := accounthandler.NewHandler(accountUsecase)

	return &Handlers{
		AccountHandler: accountHandler,
	}
}
