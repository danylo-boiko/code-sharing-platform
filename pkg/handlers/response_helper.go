package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrorType interface {
	ExecutionError | ValidationError
}

type ExecutionErrorType string

const (
	DatabaseError        ExecutionErrorType = "database_error"
	UnavailableDataError                    = "unavailable_data_error"
	IncorrectDataError                      = "incorrect_data_error"
	PermissionError                         = "permission_error"
	UnauthorizedError                       = "unauthorized"
)

type ExecutionError struct {
	ErrorType ExecutionErrorType
	Message   string
}

func NewExecutionError(errorType ExecutionErrorType, message string) ExecutionError {
	return ExecutionError{
		ErrorType: errorType,
		Message:   message,
	}
}

type Response struct {
	Success bool
	Message string
	Values  interface{}
	Errors  interface{}
}

func BadRequestResponse(context *gin.Context, message string, errors []ExecutionError) {
	context.AbortWithStatusJSON(http.StatusBadRequest, Response{
		Success: false,
		Errors:  errors,
		Message: message,
	})
}

func BadRequestValidationResponse(context *gin.Context, err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		handledErrors := GetValidationErrors(validationErrors)
		context.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Errors:  handledErrors,
		})
	}
}

func UnauthorizedResponse(context *gin.Context, message string, errors []ExecutionError) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, Response{
		Success: false,
		Errors:  errors,
		Message: message,
	})
}

func OkResponse(context *gin.Context, message string, data interface{}) {
	context.AbortWithStatusJSON(http.StatusOK, Response{
		Success: true,
		Values:  data,
		Message: message,
	})
}
