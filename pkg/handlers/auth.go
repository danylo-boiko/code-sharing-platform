package handlers

import (
	"code-sharing-platform/pkg/handlers/response"
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/requests/auth"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {
	var signInRequest auth.SignInRequest
	if err := c.ShouldBind(&signInRequest); err != nil {
		response.BadRequestValidationResponse(c, err)
		return
	}

	user, err := h.services.Authorization.GetUser(signInRequest.Username)
	if err != nil {
		executionError := response.NewExecutionError(response.IncorrectDataError, "User with this username isn't exist")
		response.UnauthorizedResponse(c, "", []response.ExecutionError{executionError})
		return
	}

	if isPasswordCorrect := h.services.Authorization.IsPasswordCorrect(signInRequest.Password, user.PasswordHash); !isPasswordCorrect {
		executionError := response.NewExecutionError(response.IncorrectDataError, "Provided wrong password")
		response.UnauthorizedResponse(c, "", []response.ExecutionError{executionError})
		return
	}

	session, err := h.services.Session.CreateSession(user.Id)
	if err != nil {
		executionError := response.NewExecutionError(response.DatabaseError, err.Error())
		response.UnauthorizedResponse(c, "", []response.ExecutionError{executionError})
		return
	}

	SaveTokenToCookie(c, session.Token, session.ExpiryDate)

	response.OkResponse(c, "User signed in successfully", nil)
}

func (h *Handler) SignUp(c *gin.Context) {
	var signUpRequest auth.SignUpRequest
	if err := c.ShouldBind(&signUpRequest); err != nil {
		response.BadRequestValidationResponse(c, err)
		return
	}

	userId, err := h.services.Authorization.CreateUser(models.User{
		Username:     signUpRequest.Username,
		Email:        signUpRequest.Email,
		PasswordHash: h.services.Authorization.HashPassword(signUpRequest.Password),
	})
	if err != nil {
		executionError := response.NewExecutionError(response.DatabaseError, err.Error())
		response.BadRequestResponse(c, "", []response.ExecutionError{executionError})
		return
	}

	session, err := h.services.Session.CreateSession(userId)
	if err != nil {
		executionError := response.NewExecutionError(response.DatabaseError, err.Error())
		response.BadRequestResponse(c, "", []response.ExecutionError{executionError})
		return
	}

	SaveTokenToCookie(c, session.Token, session.ExpiryDate)

	response.OkResponse(c, "User signed up successfully", nil)
}
