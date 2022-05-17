package services

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/repositories/interfaces"
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
		if codeSnippet.ExpiryDate.After(time.Now().UTC()) {
			return false, nil
		}
		return true, errors.New("code snippet expiry date is ended")
	}
	return false, nil
}

func (c *CodeSnippetService) IsViewsLimitReached(codeSnippet models.CodeSnippet) (bool, error) {
	if codeSnippet.ViewsLimit > 0 {
		if codeSnippet.Views < codeSnippet.ViewsLimit {
			return false, nil
		}
		return true, errors.New("views limit reached")
	}
	return false, nil
}

func (c *CodeSnippetService) AddView(codeSnippet models.CodeSnippet) error {
	codeSnippet.Views += 1
	return c.codeSnippetRepository.UpdateCodeSnippet(codeSnippet)
}

func (c *CodeSnippetService) UpdateCodeSnippet(codeSnippet models.CodeSnippet) error {
	return c.codeSnippetRepository.UpdateCodeSnippet(codeSnippet)
}

func (c *CodeSnippetService) DeleteCodeSnippet(id int) error {
	return c.codeSnippetRepository.DeleteCodeSnippet(id)
}
