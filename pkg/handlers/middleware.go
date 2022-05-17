package handlers

import (
	"code-sharing-platform/pkg/models"
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

func (h *Handler) AnonymousUserIdentity(c *gin.Context) {
	_, err := c.Cookie(codeSharingPlatformCookie)
	if err == nil {
		h.UserIdentity(c)
	}
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

func GetUserRoleClaim(userId, ownerId int) models.RoleClaimType {
	if userId == ownerId {
		return models.OwnedRoleClaim
	}
	return models.ForeignRoleClaim
}

func SaveTokenToCookie(c *gin.Context, token string, expireDate time.Time) {
	maxTokenAge := int(expireDate.Sub(time.Now().UTC()).Seconds())
	c.SetCookie(codeSharingPlatformCookie, token, maxTokenAge, "/", viper.GetString("app.domain"), false, true)
}
