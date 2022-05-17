package handlers

import (
	"code-sharing-platform/pkg/requests/auth"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {
	var signInRequest auth.SignInRequest
	if err := c.ShouldBind(&signInRequest); err != nil {
		BadRequestValidationResponse(c, err)
		return
	}

	user, err := h.services.Authorization.GetUserByUsername(signInRequest.Username)
	if err != nil {
		executionError := NewExecutionError(IncorrectDataError, "User with this username isn't exist")
		UnauthorizedResponse(c, "", []ExecutionError{executionError})
		return
	}

	if isPasswordCorrect := h.services.Authorization.IsPasswordCorrect(signInRequest.Password, user.PasswordHash); !isPasswordCorrect {
		executionError := NewExecutionError(IncorrectDataError, "Provided wrong password")
		UnauthorizedResponse(c, "", []ExecutionError{executionError})
		return
	}

	session, err := h.services.Session.CreateSession(user.Id)
	if err != nil {
		executionError := NewExecutionError(DatabaseError, err.Error())
		UnauthorizedResponse(c, "", []ExecutionError{executionError})
		return
	}

	SaveTokenToCookie(c, session.Token, session.ExpiryDate)

	OkResponse(c, "User signed in successfully", nil)
}

func (h *Handler) SignUp(c *gin.Context) {
	var signUpRequest auth.SignUpRequest
	if err := c.ShouldBind(&signUpRequest); err != nil {
		BadRequestValidationResponse(c, err)
		return
	}

	userId, err := h.services.Authorization.CreateUser(signUpRequest)
	if err != nil {
		executionError := NewExecutionError(DatabaseError, err.Error())
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	session, err := h.services.Session.CreateSession(userId)
	if err != nil {
		executionError := NewExecutionError(DatabaseError, err.Error())
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	SaveTokenToCookie(c, session.Token, session.ExpiryDate)

	OkResponse(c, "User signed up successfully", nil)
}
