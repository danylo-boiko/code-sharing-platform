package pkg

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gobeam/stringy"
)

func GetErrorMessage(fieldError validator.FieldError) string {
	field := stringy.New(fieldError.Field()).SnakeCase().ToLower()
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("'%s'is required", field)
	case "max":
		return fmt.Sprintf("'%s' should be less or equal than %s", field, fieldError.Param())
	case "min":
		return fmt.Sprintf("'%s' should be greater or equal than %s", field, fieldError.Param())
	case "email":
		return fmt.Sprintf("'%s' should be in email format", field)
	}
	return "Unknown error"
}

func GetValidationErrorsMessages(validationErrors validator.ValidationErrors) []string {
	out := make([]string, len(validationErrors))
	for i, fieldError := range validationErrors {
		out[i] = GetErrorMessage(fieldError)
	}
	return out
}
