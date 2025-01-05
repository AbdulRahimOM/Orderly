package di

import (
	accounthandler "orderly/internal/api/handler/account"
	"orderly/internal/infrastructure/config"
	repo "orderly/internal/repository"
	accountuc "orderly/internal/usecase/account"
	twilioOTP "orderly/pkg/twilio"

	"gorm.io/gorm"
)

//dependency injection

type Handlers struct {
	AccountHandler *accounthandler.Handler
}

func GetHandlers(db *gorm.DB) *Handlers {
	repo := repo.NewRepository(db)

	twilioClient := twilioOTP.NewTwilioClient(
		config.Configs.Twilio.AccountSid,
		config.Configs.Twilio.AuthToken,
		config.Configs.Twilio.ServiceSid,
		config.Configs.DevelopmentConfig.Dev_BypassOtp,
	)
	accountUsecase := accountuc.NewUsecase(repo, twilioClient)

	accountHandler := accounthandler.NewHandler(accountUsecase)

	return &Handlers{
		AccountHandler: accountHandler,
	}
}
