package repositories

import (
	"code-sharing-platform/pkg/models"
	"gorm.io/gorm"
)

type CodeSnippetMsSQL struct {
	mssql *gorm.DB
}

func NewCodeSnippetMsSQL(mssql *gorm.DB) *CodeSnippetMsSQL {
	return &CodeSnippetMsSQL{mssql: mssql}
}

func (c *CodeSnippetMsSQL) GetCodeSnippet(id int) (models.CodeSnippet, error) {
	var codeSnippet models.CodeSnippet
	if err := c.mssql.First(&codeSnippet, id).Error; err != nil {
		return codeSnippet, err
	}
	return codeSnippet, nil
}

func (s *CodeSnippetMsSQL) UpdateCodeSnippet(codeSnippet models.CodeSnippet) error {
	if err := s.mssql.Save(&codeSnippet).Error; err != nil {
		return err
	}
	return nil
}

func (c *CodeSnippetMsSQL) DeleteCodeSnippet(id int) error {
	if err := c.mssql.Delete(models.CodeSnippet{}, id).Error; err != nil {
		return err
	}
	return nil
}
