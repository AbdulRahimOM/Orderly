package accounthandler

import accountuc "orderly/internal/usecase/account"

// import usecase "orderly/internal/usecase/interface"

// type Handler struct {
// 	uc usecase.AccountUsecase
// }

// func NewHandler(uc usecase.AccountUsecase) *Handler {
// 	return &Handler{uc: uc}
// }

type Handler struct {
	uc *accountuc.AccountUC
}

func NewHandler(uc *accountuc.AccountUC) *Handler {
	return &Handler{uc: uc}
}
