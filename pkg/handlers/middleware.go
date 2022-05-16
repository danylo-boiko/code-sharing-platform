package handlers

import (
	"code-sharing-platform/pkg/handlers/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
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
		return
	}

	userId, err := h.services.Session.GetUserId(sessionToken)
	if err != nil {
		executionError := response.NewExecutionError(response.UnauthorizedError, err.Error())
		response.BadRequestResponse(c, "", []response.ExecutionError{executionError})
		return
	}

	expireDate, err := h.services.Session.ExtendExpireDate(sessionToken)
	if err != nil {
		executionError := response.NewExecutionError(response.UnauthorizedError, err.Error())
		response.BadRequestResponse(c, "", []response.ExecutionError{executionError})
		return
	}

	SaveTokenToCookie(c, sessionToken, expireDate)

	c.Set(userContext, userId)
}

func SaveTokenToCookie(c *gin.Context, token string, expireDate time.Time) {
	maxTokenAge := int(expireDate.Sub(time.Now().UTC()).Seconds())
	c.SetCookie(codeSharingPlatformCookie, token, maxTokenAge, "/", viper.GetString("app.domain"), false, true)
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
