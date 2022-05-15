package repositories

import (
	"code-sharing-platform/pkg/repositories/interfaces"
	"gorm.io/gorm"
)

type Repository struct {
	interfaces.Authorization
	interfaces.Session
}

func NewRepository(mssql *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMsSQL(mssql),
		Session:       NewSessionMsSQL(mssql),
	}
}
