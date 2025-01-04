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
type AccountUC struct {
	repo *repo.Repo
}

func NewUsecase(repo *repo.Repo) *AccountUC {
	return &AccountUC{repo: repo}
}
