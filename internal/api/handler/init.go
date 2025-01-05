package handler

import uc "orderly/internal/usecase"

// import usecase "orderly/internal/usecase/interface"

// type Handler struct {
// 	uc usecase.Usecase
// }

// func NewHandler(uc usecase.Usecase) *Handler {
// 	return &Handler{uc: uc}
// }

type Handler struct {
	uc *uc.Usecase
}

func NewHandler(uc *uc.Usecase) *Handler {
	return &Handler{uc: uc}
}
