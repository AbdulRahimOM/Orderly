package accountuc

import (
	repo "orderly/internal/repository"
	// interfaces "orderly/internal/repository/interface"
	// usecases "orderly/internal/usecase/interface"
)

// type AccountUC struct {
// 	repo interfaces.Repository
// }

//	func NewAccountUsecase(repo interfaces.Repository) usecases.AccountUsecase {
//		return &AccountUC{repo: repo}
//	}

type smsOtpClient interface {
	SendOtp(phoneNumber string) error
	VerifyOtp(phoneNumber string, otp string) (bool,error)
}

type AccountUC struct {
	repo         *repo.Repo
	smsOtpClient smsOtpClient
}

func NewUsecase(repo *repo.Repo, smsOtpClient smsOtpClient) *AccountUC {
	return &AccountUC{repo: repo, smsOtpClient: smsOtpClient}
}
