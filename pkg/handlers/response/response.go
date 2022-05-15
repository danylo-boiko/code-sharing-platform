package response

import (
	"code-sharing-platform/pkg"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrorType interface {
	ExecutionError | pkg.ValidationError
}

type ExecutionErrorType string

const (
	DatabaseError      ExecutionErrorType = "database_error"
	IncorrectDataError                    = "incorrect_data_error"
	UnauthorizedError                     = "unauthorized"
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

type Response[T ErrorType] struct {
	Success bool
	Message string
	Values  interface{}
	Errors  []T
}

func BadRequestResponse(context *gin.Context, message string, errors []ExecutionError) {
	context.AbortWithStatusJSON(http.StatusBadRequest, Response[ExecutionError]{
		Success: false,
		Errors:  errors,
		Message: message,
	})
}

func BadRequestValidationResponse(context *gin.Context, err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		handledErrors := pkg.GetValidationErrors(validationErrors)
		context.AbortWithStatusJSON(http.StatusBadRequest, Response[pkg.ValidationError]{
			Success: false,
			Errors:  handledErrors,
		})
	}
}

func UnauthorizedResponse(context *gin.Context, message string, errors []ExecutionError) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, Response[ExecutionError]{
		Success: false,
		Errors:  errors,
		Message: message,
	})
}

func OkResponse(context *gin.Context, message string, data interface{}) {
	context.AbortWithStatusJSON(http.StatusOK, Response[ExecutionError]{
		Success: true,
		Values:  data,
		Message: message,
	})
}
