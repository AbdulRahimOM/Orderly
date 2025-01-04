package accounthandler

import usecase "orderly/internal/usecase/interface"

type Handler struct {
	uc usecase.AccountUsecase
}

func NewHandler(uc usecase.AccountUsecase) *Handler {
	return &Handler{uc: uc}
}
