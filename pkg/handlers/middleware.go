package handlers

import (
	"code-sharing-platform/pkg/handlers/response"
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	codeSharingPlatformCookie = "code_sharing_platform"
	userContext               = "user_id"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	sessionToken, err := c.Cookie(codeSharingPlatformCookie)
	if err != nil {
		executionError := response.NewExecutionError(response.UnauthorizedError, "Session token isn't exist")
		response.BadRequestResponse(c, "", []response.ExecutionError{executionError})
	}

	userId, err := h.services.Session.GetUserId(sessionToken)
	if err != nil {
		executionError := response.NewExecutionError(response.DatabaseError, err.Error())
		response.BadRequestResponse(c, "", []response.ExecutionError{executionError})
	}

	c.Set(userContext, userId)
}

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userContext)
	if !ok {
		return 0, errors.New("user id isn't exist")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id has invalid type")
	}

	return idInt, nil
}
