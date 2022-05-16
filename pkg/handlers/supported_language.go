package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetSupportedLanguages(c *gin.Context) {
	supportedLanguages, err := h.services.SupportedLanguage.GetSupportedLanguages()
	if err != nil {
		executionError := NewExecutionError(DatabaseError, err.Error())
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	OkResponse(c, "", supportedLanguages)
}
