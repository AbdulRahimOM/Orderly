package di

import (
	handler "orderly/internal/api/handler"
	"orderly/internal/infrastructure/config"
	repo "orderly/internal/repository"
	uc "orderly/internal/usecase"
	twilioOTP "orderly/pkg/twilio"

	"gorm.io/gorm"
)

//dependency injection

type Handlers struct {
	Handler *handler.Handler
}

func GetHandlers(db *gorm.DB) *Handlers {
	repo := repo.NewRepository(db)

	twilioClient := twilioOTP.NewTwilioClient(
		config.Configs.Twilio.AccountSid,
		config.Configs.Twilio.AuthToken,
		config.Configs.Twilio.ServiceSid,
		config.Configs.DevelopmentConfig.Dev_BypassOtp,
	)
	accountUsecase := uc.NewUsecase(repo, twilioClient)

	accountHandler := handler.NewHandler(accountUsecase)

	return &Handlers{
		Handler: accountHandler,
	}
}
