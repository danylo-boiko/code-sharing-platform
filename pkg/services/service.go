package services

import (
	"code-sharing-platform/pkg/repositories"
	"code-sharing-platform/pkg/services/interfaces"
)

type Service struct {
	interfaces.Authorization
	interfaces.Session
	interfaces.Role
	interfaces.CodeSnippet
	interfaces.SupportedLanguage
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		Authorization:     NewAuthService(repository.Authorization, repository.Role),
		Session:           NewSessionService(repository.Session),
		Role:              NewRoleService(repository.Role),
		CodeSnippet:       NewCodeSnippetService(repository.CodeSnippet),
		SupportedLanguage: NewSupportedLanguageService(repository.SupportedLanguage),
	}
}
