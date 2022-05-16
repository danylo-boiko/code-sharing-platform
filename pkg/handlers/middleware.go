package handlers

import (
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
		executionError := NewExecutionError(UnauthorizedError, "Session token isn't exist")
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	session, err := h.services.Session.GetSession(sessionToken)
	if err != nil {
		executionError := NewExecutionError(UnauthorizedError, "No session with such token")
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	if session.ExpiryDate.Before(time.Now().UTC()) {
		executionError := NewExecutionError(UnauthorizedError, "Session token no longer valid, sign in again")
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	expireDate, err := h.services.Session.ExtendExpireDate(sessionToken)
	if err != nil {
		executionError := NewExecutionError(UnauthorizedError, err.Error())
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	SaveTokenToCookie(c, sessionToken, expireDate)

	c.Set(userContext, session.UserId)
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
