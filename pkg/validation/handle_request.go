package validation

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"orderly/internal/domain/response"

	"github.com/gofiber/fiber/v2"
)

const (
	bindingErrCode      = "BINDING_ERROR"
	validationErrCode   = "VALIDATION_ERROR"
	queryBindingErrCode = "URL QUERY BINDING ERROR"
)

// BindAndValidateRequest binds and validates the request.
// Req should be a pointer to the request struct.
func BindAndValidateJSONRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := c.BodyParser(req); err != nil {
		log.Println("error parsing request:", err)
		return false, response.Response{
			HttpStatusCode: http.StatusBadRequest,
			Status:         false,
			ResponseCode:   bindingErrCode,
			Error:          err,
		}.WriteToJSON(c)
	}
	if err := validateJSONRequestDetailed(req); err != nil {
		log.Println("error validating request:", err)
		return false, c.Status(http.StatusBadRequest).JSON(response.ValidationErrorResponse{
			Status:       false,
			ResponseCode: validationErrCode,
			Errors:       err,
		})
	}

	fmt.Println("req after validation:", req)
	return true, nil
}

// Req should be a pointer to the request struct.
func BindAndValidateArrayJSONRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := c.BodyParser(req); err != nil {
		log.Println("error parsing request:", err)
		return false, response.Response{
			HttpStatusCode: http.StatusBadRequest,
			Status:         false,
			ResponseCode:   bindingErrCode,
			Error:          err,
		}.WriteToJSON(c)
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

	fmt.Println("req after validation:", req)
	return true, nil
}

func BindAndValidateURLQueryRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := c.QueryParser(req); err != nil {
		log.Println("error parsing request:", err)
		return false, response.Response{
			HttpStatusCode: http.StatusBadRequest,
			Status:         false,
			ResponseCode:   queryBindingErrCode,
			Error:          err,
		}.WriteToJSON(c)
	}
	if err := validateJSONRequestDetailed(req); err != nil {
		log.Println("error validating request:", err)
		return false, c.Status(http.StatusBadRequest).JSON(response.ValidationErrorResponse{
			Status:       false,
			ResponseCode: validationErrCode,
			Errors:       err,
		})
	}

	fmt.Println("req after validation:", req)
	return true, nil
}

// BindAndValidateFormDataRequest seems to work only for form data requests with string valued fields.
// Req should be a pointer to the request struct.
func BindAndValidateFormDataRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := c.BodyParser(req); err != nil {
		log.Println("error parsing request:", err)
		return false, response.Response{
			Status:       false,
			ResponseCode: bindingErrCode,
			Error:        err,
		}.WriteToJSON(c)
	}

	if err := validateFormDataRequestDetailed(req); err != nil {
		log.Println("error validating request:", err)
		return false, c.Status(http.StatusBadRequest).JSON(response.ValidationErrorResponse{
			Status:       false,
			ResponseCode: validationErrCode,
			Errors:       err,
		})
	}

	fmt.Println("req after validation:", req)

	return true, nil
}

// Validate form data request
func ValidateFormDataRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := validateFormDataRequestDetailed(req); err != nil {
		log.Println("error validating request:", err)
		return false, c.Status(http.StatusBadRequest).JSON(response.ValidationErrorResponse{
			Status:       false,
			ResponseCode: validationErrCode,
			Errors:       err,
		})
	}

	fmt.Println("req after validation:", req)

	return true, nil
}

// //Validate form data request
// func ValidateFormDataRequest(c *fiber.Ctx, req interface{}) (bool, error) {
// 	if err := validateFormDataRequestDetailed(req); err != nil {
// 		log.Println("error validating request:", err)
// 		return false, c.Status(http.StatusBadRequest).JSON(response.ValidationErrorResponse{
// 			Status:       false,
// 			ResponseCode: validationErrCode,
// 			Errors:       err,
// 		})
// 	}

// 	return true, nil
// }
