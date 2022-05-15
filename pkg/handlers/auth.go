package handlers

import (
	"code-sharing-platform/pkg/handlers/response"
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/requests/auth"
	"github.com/gin-gonic/gin"
	"time"
)

func (h *Handler) SignIn(c *gin.Context) {
	var signInRequest auth.SignInRequest
	if err := c.ShouldBind(&signInRequest); err != nil {
		response.BadRequestValidationErrors(c, err)
		return
	}

	sessionToken, expireDate, err := h.services.Authorization.CreateSessionToken(signInRequest.Username, signInRequest.Password)

	if err != nil {
		response.BadRequest(c, "Error occurred while creating new session token", []response.ErrorDetail{{
			ErrorType:    response.DatabaseError,
			ErrorMessage: err.Error(),
		}})
		return
	}

	maxTokenAge := int(expireDate.Sub(time.Now().UTC()).Seconds())

	c.SetCookie("code_sharing_platform", sessionToken, maxTokenAge, "/", "localhost", false, true)

	response.OkRequest(c, "User signed in successfully", nil)
}

func (h *Handler) SignUp(c *gin.Context) {
	var signUpRequest auth.SignUpRequest
	if err := c.ShouldBind(&signUpRequest); err != nil {
		response.BadRequestValidationErrors(c, err)
		return
	}

	_, err := h.services.Authorization.CreateUser(models.User{
		Username:     signUpRequest.Username,
		Email:        signUpRequest.Email,
		PasswordHash: h.services.Authorization.HashPassword(signUpRequest.Password),
	})
	if err != nil {
		response.BadRequest(c, "Error occurred while creating new user", []response.ErrorDetail{{
			ErrorType:    response.DatabaseError,
			ErrorMessage: err.Error(),
		}})
		return
	}

	sessionToken, expireDate, err := h.services.Authorization.CreateSessionToken(signUpRequest.Username, signUpRequest.Password)
	if err != nil {
		response.BadRequest(c, "Error occurred while creating new session token", []response.ErrorDetail{{
			ErrorType:    response.DatabaseError,
			ErrorMessage: err.Error(),
		}})
		return
	}

	maxTokenAge := int(expireDate.Sub(time.Now().UTC()).Seconds())

	c.SetCookie("code_sharing_platform", sessionToken, maxTokenAge, "/", "localhost", false, true)

	response.OkRequest(c, "User signed up successfully", nil)
}
