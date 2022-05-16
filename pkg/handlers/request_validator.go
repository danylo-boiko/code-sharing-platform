package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gobeam/stringy"
)

type ValidationError struct {
	Field        string
	ErrorMessage string
}

func GetErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	case "max":
		return "Should be less or equal than " + fieldError.Param()
	case "min":
		return "Should be greater or equal than " + fieldError.Param()
	case "email":
		return "Should be in email format"
	}
	return "Unknown error"
}

func GetValidationErrors(validationErrors validator.ValidationErrors) []ValidationError {
	out := make([]ValidationError, len(validationErrors))
	for idx, fieldError := range validationErrors {
		out[idx] = ValidationError{
			Field:        stringy.New(fieldError.Field()).SnakeCase().ToLower(),
			ErrorMessage: GetErrorMessage(fieldError),
		}
	}
	return out
}
