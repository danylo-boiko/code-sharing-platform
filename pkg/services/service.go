package services

import (
	"code-sharing-platform/pkg/repositories"
	"code-sharing-platform/pkg/services/interfaces"
)

type Service struct {
	interfaces.Authorization
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization, repository.Session),
	}
}
