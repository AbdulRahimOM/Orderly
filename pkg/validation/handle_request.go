package validation

import (
	"fmt"
	"log"
	"net/http"
	"orderly/internal/domain/response"
	"orderly/internal/infrastructure/config"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

const (
	bindingErrCode      = "BINDING_ERROR"
	validationErrCode   = "VALIDATION_ERROR"
	queryBindingErrCode = "URL_QUERY_BINDING_ERROR"
)

type customValidation interface {
	CustomValidation() (responseCode string, err error)
}

func bindErrResponse(c *fiber.Ctx, err error) (bool, error) {
	log.Println("error parsing request:", err)
	return false, response.Response{
		HttpStatusCode: http.StatusBadRequest,
		Status:         false,
		ResponseCode:   bindingErrCode,
		Error:          err,
	}.WriteToJSON(c)
}

func validationErrResponse(c *fiber.Ctx, err []response.InvalidField) (bool, error) {
	log.Println("error validating request:", err)
	return false, c.Status(http.StatusBadRequest).JSON(response.ValidationErrorResponse{
		Status:       false,
		ResponseCode: validationErrCode,
		Errors:       err,
	})
}

func customValidationErrResponse(c *fiber.Ctx, responseCode string, err error) (bool, error) {
	return false, response.Response{
		HttpStatusCode: http.StatusBadRequest,
		Status:         false,
		ResponseCode:   responseCode,
		Error:          err,
	}.WriteToJSON(c)
}

// BindAndValidateRequest binds and validates the request.
// Req should be a pointer to the request struct.
func BindAndValidateJSONRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := c.BodyParser(req); err != nil {
		return bindErrResponse(c, err)
	}
	if err := validateJSONRequestDetailed(req); err != nil {
		return validationErrResponse(c, err)
	}

	if _, ok := req.(customValidation); ok {
		if responseCode, err := req.(customValidation).CustomValidation(); err != nil {
			return customValidationErrResponse(c, responseCode, err)
		}
	}

	if config.Configs.Dev_Mode {
		fmt.Println("#Dev: Req after validation:", req)
	}
	return true, nil
}

// Req should be a pointer to the request struct.
func BindAndValidateArrayJSONRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := c.BodyParser(req); err != nil {
		return bindErrResponse(c, err)
	}

	val := reflect.ValueOf(req)
	val = val.Elem()
	log.Println(val.Kind())
	if val.Kind() != reflect.Slice {
		return false, c.Status(http.StatusBadRequest).JSON(response.ValidationErrorResponse{
			Status:       false,
			ResponseCode: validationErrCode,
		})
	}

	for i := 0; i < val.Len(); i++ {
		item := val.Index(i).Interface()
		if err := validateJSONRequestDetailed(item); err != nil {
			log.Println("Error validating request:", err)
			return false, c.Status(http.StatusBadRequest).JSON(response.ValidationErrorResponse{
				Status:       false,
				ResponseCode: validationErrCode,
				Errors:       err,
			})
		}
	}

	if _, ok := req.(customValidation); ok {
		if responseCode, err := req.(customValidation).CustomValidation(); err != nil {
			return customValidationErrResponse(c, responseCode, err)
		}
	}

	if config.Configs.Dev_Mode {
		fmt.Println("#Dev: Req after validation:", req)
	}
	return true, nil
}

func BindAndValidateURLQueryRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := c.QueryParser(req); err != nil {
		return bindErrResponse(c, err)
	}
	if err := validateJSONRequestDetailed(req); err != nil {
		return validationErrResponse(c, err)
	}

	if _, ok := req.(customValidation); ok {
		if responseCode, err := req.(customValidation).CustomValidation(); err != nil {
			return customValidationErrResponse(c, responseCode, err)
		}
	}

	if config.Configs.Dev_Mode {
		fmt.Println("#Dev: Req after validation:", req)
	}
	return true, nil
}

// BindAndValidateFormDataRequest seems to work only for form data requests with string valued fields.
// Req should be a pointer to the request struct.
func BindAndValidateFormDataRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := c.BodyParser(req); err != nil {
		return bindErrResponse(c, err)
	}

	if err := validateFormDataRequestDetailed(req); err != nil {
		return validationErrResponse(c, err)
	}

	if _, ok := req.(customValidation); ok {
		if responseCode, err := req.(customValidation).CustomValidation(); err != nil {
			return customValidationErrResponse(c, responseCode, err)
		}
	}

	if config.Configs.Dev_Mode {
		fmt.Println("#Dev: Req after validation:", req)
	}

	return true, nil
}

// Validate form data request
func ValidateFormDataRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := validateFormDataRequestDetailed(req); err != nil {
		return validationErrResponse(c, err)
	}

	if _, ok := req.(customValidation); ok {
		if responseCode, err := req.(customValidation).CustomValidation(); err != nil {
			return customValidationErrResponse(c, responseCode, err)
		}
	}

	if config.Configs.Dev_Mode {
		fmt.Println("#Dev: Req after validation:", req)
	}

	return true, nil
}
