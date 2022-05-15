package response

import (
	"code-sharing-platform/pkg"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrorType string

const (
	Error             ErrorType = "Error"
	DatabaseError     ErrorType = "Database error"
	ValidationError   ErrorType = "Validation error"
	UnauthorizedError ErrorType = "Unauthorized"
)

type ErrorDetail struct {
	ErrorType    ErrorType
	ErrorMessage string
}

type Response struct {
	Success bool
	Value   []interface{}
	Message string
	Errors  []ErrorDetail
}

func BadRequest(context *gin.Context, message string, errors []ErrorDetail) {
	context.AbortWithStatusJSON(http.StatusBadRequest, Response{
		Success: false,
		Errors:  errors,
		Message: message,
	})
}

func BadRequestValidationErrors(context *gin.Context, err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		handledErrors := make([]ErrorDetail, len(validationErrors))
		for i, message := range pkg.GetValidationErrorsMessages(validationErrors) {
			handledErrors[i] = ErrorDetail{
				ErrorType:    ValidationError,
				ErrorMessage: message,
			}
		}
		BadRequest(context, "Fields validation errors", handledErrors)
	}
}

func OkRequest(context *gin.Context, message string, data []interface{}) {
	context.AbortWithStatusJSON(http.StatusOK, Response{
		Success: true,
		Value:   data,
		Message: message,
	})
}
