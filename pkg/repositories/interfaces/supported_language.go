package interfaces

import "code-sharing-platform/pkg/models"

type SupportedLanguage interface {
	GetSupportedLanguages() ([]models.SupportedLanguage, error)
}
