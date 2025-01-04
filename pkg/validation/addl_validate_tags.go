package validation

import (
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func init() {
	validate.RegisterValidation("pincode", validatePincode) //=> we can use 'pincode' as a tag in the struct field to validate the pincode
	validate.RegisterValidation("not_num_only", notNumOnly)
	validate.RegisterValidation("alpha_space_dot", isAlphaSpaceDot)
	validate.RegisterValidation("contains_alphabet", containsAlphabet)
}

func validatePincode(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if len(value) != 6 {
		return false
	}
	pinCodeNum, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	//check if decimal
	if pinCodeNum != float64(int64(pinCodeNum)) {
		return false
	}
	return (pinCodeNum >= 110000 && pinCodeNum <= 899999)
}

func notNumOnly(fl validator.FieldLevel) bool {
	// Regular expression to check if the string contains only digits
	isNumOnly := regexp.MustCompile(`^\d+$`).MatchString
	// Return true if the string is not purely numeric
	return !isNumOnly(fl.Field().String())
}

func isAlphaSpaceDot(fl validator.FieldLevel) bool {
	// Regular expression to check if the string contains only alphabets, spaces, and dots
	isValid := regexp.MustCompile(`^[a-zA-Z\s.]+$`).MatchString
	// Return true if the string matches the allowed pattern
	return isValid(fl.Field().String())
}

func containsAlphabet(fl validator.FieldLevel) bool {
    // Regular expression to check if the string contains at least one alphabet character
    hasAlphabet := regexp.MustCompile(`[A-Za-z]`).MatchString
    // Return true if the string contains at least one alphabet character
    return hasAlphabet(fl.Field().String())
}