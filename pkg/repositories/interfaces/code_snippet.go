package interfaces

import "code-sharing-platform/pkg/models"

type CodeSnippet interface {
	GetCodeSnippet(id int) (models.CodeSnippet, error)
	UpdateCodeSnippet(codeSnippet models.CodeSnippet) error
	DeleteCodeSnippet(id int) error
}
