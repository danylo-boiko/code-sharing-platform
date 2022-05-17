package interfaces

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/requests/code_snippet"
)

type CodeSnippet interface {
	GetCodeSnippet(id int) (models.CodeSnippet, error)
	IsExpiryDateEnded(codeSnippet models.CodeSnippet) (bool, error)
	IsViewsLimitReached(codeSnippet models.CodeSnippet) (bool, error)
	AddView(codeSnippet models.CodeSnippet) error
	CreateCodeSnippet(userId int, request code_snippet.CreateCodeSnippetRequest) (int, error)
	UpdateCodeSnippet(id int, codeSnippet code_snippet.UpdateCodeSnippetRequest) error
	DeleteCodeSnippet(id int) error
}
