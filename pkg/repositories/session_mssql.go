package repositories

import (
	"code-sharing-platform/pkg/models"
	"gorm.io/gorm"
)

type SessionMsSQL struct {
	mssql *gorm.DB
}

func NewSessionMsSQL(mssql *gorm.DB) *SessionMsSQL {
	return &SessionMsSQL{mssql: mssql}
}

func (s *SessionMsSQL) GetSession(token string) (models.Session, error) {
	var session models.Session
	if err := s.mssql.First(&session, "token = ?", token).Error; err != nil {
		return session, err
	}
	return session, nil
}

func (s *SessionMsSQL) CreateSession(session models.Session) (int, error) {
	if err := s.mssql.Create(&session).Error; err != nil {
		return session.Id, err
	}
	return session.Id, nil
}
