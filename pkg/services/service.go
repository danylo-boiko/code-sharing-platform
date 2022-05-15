package services

import (
	"code-sharing-platform/pkg/repositories"
	"code-sharing-platform/pkg/services/interfaces"
)

type Service struct {
	interfaces.Authorization
	interfaces.Session
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization, repository.Session),
		Session:       NewSessionService(repository.Session),
	}
}
