package handlers

import (
	"code-sharing-platform/pkg/handlers/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetSupportedLanguages(c *gin.Context) {
	supportedLanguages, err := h.services.SupportedLanguage.GetSupportedLanguages()
	if err != nil {
		executionError := response.NewExecutionError(response.DatabaseError, err.Error())
		response.BadRequestResponse(c, "", []response.ExecutionError{executionError})
		return
	}

	response.OkResponse(c, "", supportedLanguages)
}
