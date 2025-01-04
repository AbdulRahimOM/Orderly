package accountuc

import (
	interfaces "orderly/internal/repository/interface"
	usecases "orderly/internal/usecase/interface"
)

type AccountUC struct {
	repo interfaces.Repository
}

func NewAccountUsecase(repo interfaces.Repository) usecases.AccountUsecase {
	return &AccountUC{repo: repo}
}
