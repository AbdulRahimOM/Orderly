package uc

import (
	interfaces "orderly/internal/repository/interface"
	usecases "orderly/internal/usecase/interface"
)

type Usecase struct {
	repo         interfaces.Repository
	smsOtpClient smsOtpClient
}

type smsOtpClient interface {
	SendOtp(phoneNumber string) error
	VerifyOtp(phoneNumber string, otp string) (bool, error)
}

func NewUsecase(repo interfaces.Repository, smsOtpClient smsOtpClient) usecases.Usecase {
	return &Usecase{repo: repo, smsOtpClient: smsOtpClient}
}
