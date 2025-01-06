package handler

import (
	"orderly/internal/domain/request"
	"orderly/internal/domain/response"
	"orderly/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateCategory(c *fiber.Ctx) error {

	req := new(request.CategoryReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.CreateCategory(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) GetCategories(c *fiber.Ctx) error {
	req, errResponse := request.GetListRequest(c)
	if req == nil {
		return errResponse
	}

	response := h.uc.GetCategories(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) GetCategoryByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
	}
	response := h.uc.GetCategoryByID(c.Context(), id)
	return response.WriteToJSON(c)
}

func (h *Handler) UpdateCategoryByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
	}
	req := new(request.CategoryReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.UpdateCategoryByID(c.Context(), id, req)
	return response.WriteToJSON(c)
}

func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	req := new(request.ProductReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.CreateProduct(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) GetProducts(c *fiber.Ctx) error {
	req, errResponse := request.GetListRequest(c)
	if req == nil {
		return errResponse
	}

	response := h.uc.GetProducts(c.Context(), req)
	return response.WriteToJSON(c)
}

func (h *Handler) GetProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
	}
	response := h.uc.GetProductByID(c.Context(), id)
	return response.WriteToJSON(c)
}

func (h *Handler) UpdateProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
	}
	req := new(request.UpdateProductReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.UpdateProductByID(c.Context(), id, req)
	return response.WriteToJSON(c)
}

func (h *Handler) GetProductStockByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
	}
	response := h.uc.GetProductStockByID(c.Context(), id)
	return response.WriteToJSON(c)
}

func (h *Handler) AddProductStockByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.InvalidURLParamResponse("id", err).WriteToJSON(c)
	}
	req := new(request.AddProductStockReq)
	if ok, err := validation.BindAndValidateJSONRequest(c, req); !ok {
		return err
	}

	response := h.uc.AddProductStockByID(c.Context(), id, req)
	return response.WriteToJSON(c)
}
