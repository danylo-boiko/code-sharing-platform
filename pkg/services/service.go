package services

import "code-sharing-platform/pkg/repositories"

type Service struct {
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{}
}
