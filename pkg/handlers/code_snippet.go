package handlers

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/requests/code_snippet"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// @Tags Code Snippets
// @Param id path int true "Required"
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 400,401 {object} Response
// @Router /api/code-snippets/{id} [get]
func (h *Handler) GetCodeSnippetById(c *gin.Context) {
	codeSnippetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		executionError := NewExecutionError(IncorrectDataError, "Invalid code snippet id param")
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	codeSnippet, err := h.services.CodeSnippet.GetCodeSnippet(codeSnippetId)
	if err != nil {
		executionError := NewExecutionError(IncorrectDataError, "Code snippet with this id isn't exist")
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	userId, err := GetUserId(c)
	if err == nil {
		permissionErr := h.services.Role.CheckUserPermission(userId, models.ForeignRoleClaim, models.ReadAction)
		if codeSnippet.UserId == userId || permissionErr == nil {
			OkResponse(c, "", codeSnippet)
			return
		}
	}

	if _, err = h.services.CodeSnippet.IsExpiryDateEnded(codeSnippet); err != nil {
		executionError := NewExecutionError(UnavailableDataError, err.Error())
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	if _, err = h.services.CodeSnippet.IsViewsLimitReached(codeSnippet); err != nil {
		executionError := NewExecutionError(UnavailableDataError, err.Error())
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	if err = h.services.CodeSnippet.AddView(codeSnippet); err != nil {
		executionError := NewExecutionError(DatabaseError, err.Error())
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	OkResponse(c, "", codeSnippet)
}

// @Tags Code Snippets
// @Accept json
// @Produce json
// @Param input body code_snippet.CreateCodeSnippetRequest true "Request boby"
// @Success 200 {object} Response
// @Failure 400,401 {object} Response
// @Router /api/code-snippets [post]
func (h *Handler) CreateCodeSnippet(c *gin.Context) {
	var createCodeSnippetRequest code_snippet.CreateCodeSnippetRequest
	if err := c.ShouldBind(&createCodeSnippetRequest); err != nil {
		switch err.(type) {
		case *time.ParseError:
			executionError := NewExecutionError(IncorrectDataError, "Dates should using format '2006-01-02T15:04:05Z'")
			BadRequestResponse(c, "", []ExecutionError{executionError})
		default:
			BadRequestValidationResponse(c, err)
		}
		return
	}

	userId, _ := GetUserId(c)
	permissionErr := h.services.Role.CheckUserPermission(userId, models.OwnedRoleClaim, models.CreateAction)
	if permissionErr != nil {
		executionError := NewExecutionError(PermissionError, permissionErr.Error())
		UnauthorizedResponse(c, "", []ExecutionError{executionError})
		return
	}

	codeSnippetId, err := h.services.CodeSnippet.CreateCodeSnippet(userId, createCodeSnippetRequest)
	if err != nil {
		executionError := NewExecutionError(DatabaseError, err.Error())
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	OkResponse(c, "Code snippet created successfully", codeSnippetId)
}

// @Tags Code Snippets
// @Param id path int true "Required"
// @Param input body code_snippet.UpdateCodeSnippetRequest true "Request boby"
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 400,401 {object} Response
// @Router /api/code-snippets/{id} [put]
func (h *Handler) UpdateCodeSnippet(c *gin.Context) {
	codeSnippetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		executionError := NewExecutionError(IncorrectDataError, "Invalid code snippet id param")
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	var updateCodeSnippetRequest code_snippet.UpdateCodeSnippetRequest
	if err := c.ShouldBind(&updateCodeSnippetRequest); err != nil {
		switch err.(type) {
		case *time.ParseError:
			executionError := NewExecutionError(IncorrectDataError, "Dates should using format '2006-01-02T15:04:05Z'")
			BadRequestResponse(c, "", []ExecutionError{executionError})
		default:
			BadRequestValidationResponse(c, err)
		}
	}

	codeSnippet, err := h.services.CodeSnippet.GetCodeSnippet(codeSnippetId)
	if err != nil {
		executionError := NewExecutionError(IncorrectDataError, "Code snippet with this id isn't exist")
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	userId, _ := GetUserId(c)
	permissionErr := h.services.Role.CheckUserPermission(userId, GetUserRoleClaim(userId, codeSnippet.UserId), models.DeleteAction)
	if permissionErr != nil {
		executionError := NewExecutionError(PermissionError, permissionErr.Error())
		UnauthorizedResponse(c, "", []ExecutionError{executionError})
		return
	}

	if err := h.services.CodeSnippet.UpdateCodeSnippet(codeSnippet.Id, updateCodeSnippetRequest); err != nil {
		executionError := NewExecutionError(DatabaseError, err.Error())
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	OkResponse(c, "Code snippet updated successfully", nil)
}

// @Tags Code Snippets
// @Param id path int true "Required"
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 400,401 {object} Response
// @Router /api/code-snippets/{id} [delete]
func (h *Handler) DeleteCodeSnippet(c *gin.Context) {
	codeSnippetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		executionError := NewExecutionError(IncorrectDataError, "Invalid code snippet id param")
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	codeSnippet, err := h.services.CodeSnippet.GetCodeSnippet(codeSnippetId)
	if err != nil {
		executionError := NewExecutionError(DatabaseError, "Code snippet with this id isn't exist")
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	userId, _ := GetUserId(c)
	permissionErr := h.services.Role.CheckUserPermission(userId, GetUserRoleClaim(userId, codeSnippet.UserId), models.DeleteAction)
	if permissionErr != nil {
		executionError := NewExecutionError(PermissionError, permissionErr.Error())
		UnauthorizedResponse(c, "", []ExecutionError{executionError})
		return
	}

	err = h.services.CodeSnippet.DeleteCodeSnippet(codeSnippet.Id)
	if err != nil {
		executionError := NewExecutionError(DatabaseError, err.Error())
		BadRequestResponse(c, "", []ExecutionError{executionError})
		return
	}

	OkResponse(c, "Code snippet deleted successfully", nil)
}
