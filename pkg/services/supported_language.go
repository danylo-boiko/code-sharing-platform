package services

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/repositories/interfaces"
)

type SupportedLanguageService struct {
	languagesRepository interfaces.SupportedLanguage
}

func NewSupportedLanguageService(languagesRepository interfaces.SupportedLanguage) *SupportedLanguageService {
	return &SupportedLanguageService{languagesRepository: languagesRepository}
}

func (s *SupportedLanguageService) GetSupportedLanguages() ([]models.SupportedLanguage, error) {
	return s.languagesRepository.GetSupportedLanguages()
}
