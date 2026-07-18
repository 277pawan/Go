package validation

import (
	"fmt"

	"github.com/go-playground/validator"
)

var Validate = validator.New()

func FormatValidationErrors(err error) []map[string]string {
	var errors []map[string]string

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return []map[string]string{
			{"message": "invalid request"},
		}
	}

	for _, e := range validationErrors {
		message := ""

		switch e.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", e.Field())
		case "email":
			message = fmt.Sprintf("%s must be a valid email", e.Field())
		case "min":
			message = fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
		case "max":
			message = fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
		default:
			message = fmt.Sprintf("%s is invalid", e.Field())
		}

		errors = append(errors, map[string]string{
			"field":   e.Field(),
			"message": message,
		})
	}

	return errors
}
