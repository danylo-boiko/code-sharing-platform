package services

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/repositories/interfaces"
	"code-sharing-platform/pkg/requests/code_snippet"
	"errors"
	"time"
)

type CodeSnippetService struct {
	codeSnippetRepository interfaces.CodeSnippet
}

func NewCodeSnippetService(codeSnippetRepository interfaces.CodeSnippet) *CodeSnippetService {
	return &CodeSnippetService{codeSnippetRepository: codeSnippetRepository}
}

func (c *CodeSnippetService) GetCodeSnippet(id int) (models.CodeSnippet, error) {
	return c.codeSnippetRepository.GetCodeSnippet(id)
}

func (c *CodeSnippetService) IsExpiryDateEnded(codeSnippet models.CodeSnippet) (bool, error) {
	if !codeSnippet.ExpiryDate.IsZero() {
		if codeSnippet.ExpiryDate.Before(time.Now().UTC()) {
			return true, errors.New("code snippet expiry date is ended")
		}
	}
	return false, nil
}

func (c *CodeSnippetService) IsViewsLimitReached(codeSnippet models.CodeSnippet) (bool, error) {
	if codeSnippet.ViewsLimit > 0 {
		if codeSnippet.Views == codeSnippet.ViewsLimit {
			return true, errors.New("views limit reached")
		}
	}
	return false, nil
}

func (c *CodeSnippetService) AddView(codeSnippet models.CodeSnippet) error {
	codeSnippet.Views += 1
	return c.codeSnippetRepository.UpdateCodeSnippet(codeSnippet)
}

func (c *CodeSnippetService) CreateCodeSnippet(userId int, request code_snippet.CreateCodeSnippetRequest) (int, error) {
	var codeSnippet models.CodeSnippet

	codeSnippet.UserId = userId
	codeSnippet.LanguageId = request.LanguageId
	codeSnippet.Header = request.Header
	codeSnippet.Code = request.Code
	if !request.ExpiryDate.IsZero() {
		codeSnippet.ExpiryDate = request.ExpiryDate.UTC()
	}
	if request.ViewsLimit != 0 {
		codeSnippet.ViewsLimit = request.ViewsLimit
	}
	codeSnippet.CreatedAt = time.Now().UTC()

	codeSnippetId, err := c.codeSnippetRepository.CreateCodeSnippet(codeSnippet)
	if err != nil {
		return 0, err
	}
	return codeSnippetId, nil
}

func (c *CodeSnippetService) UpdateCodeSnippet(id int, request code_snippet.UpdateCodeSnippetRequest) error {
	codeSnippet, err := c.codeSnippetRepository.GetCodeSnippet(id)
	if err != nil {
		return err
	}

	if request.LanguageId != 0 {
		codeSnippet.LanguageId = request.LanguageId
	}
	if len(request.Header) != 0 {
		codeSnippet.Header = request.Header
	}
	if len(request.Code) != 0 {
		codeSnippet.Code = request.Code
	}
	if !request.ExpiryDate.IsZero() {
		codeSnippet.ExpiryDate = request.ExpiryDate.UTC()
	}
	if request.ViewsLimit != 0 {
		codeSnippet.ViewsLimit = request.ViewsLimit
	}

	if err = c.codeSnippetRepository.UpdateCodeSnippet(codeSnippet); err != nil {
		return err
	}

	return nil
}

func (c *CodeSnippetService) DeleteCodeSnippet(id int) error {
	return c.codeSnippetRepository.DeleteCodeSnippet(id)
}
