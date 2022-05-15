package repositories

import (
	"code-sharing-platform/pkg/models"
	"gorm.io/gorm"
)

type SupportedLanguageMsSQL struct {
	mssql *gorm.DB
}

func NewSupportedLanguageMsSQL(mssql *gorm.DB) *SupportedLanguageMsSQL {
	return &SupportedLanguageMsSQL{mssql: mssql}
}

func (s *SupportedLanguageMsSQL) GetSupportedLanguages() ([]models.SupportedLanguage, error) {
	var supportedLanguages []models.SupportedLanguage
	if err := s.mssql.Find(&supportedLanguages).Error; err != nil {
		return supportedLanguages, err
	}
	return supportedLanguages, nil
}
