package services

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/repositories/interfaces"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"time"
)

type SessionService struct {
	sessionRepository interfaces.Session
}

func NewSessionService(sessionRepository interfaces.Session) *SessionService {
	return &SessionService{sessionRepository: sessionRepository}
}

func (s *SessionService) GetUserId(sessionToken string) (int, error) {
	session, err := s.sessionRepository.GetSession(sessionToken)
	if err != nil {
		return 0, err
	}
	return session.UserId, nil
}

func (s *SessionService) CreateSession(userId int) (models.Session, error) {
	sessionToken := uuid.NewString()
	expireDate := time.Now().UTC().Add(time.Duration(viper.GetInt("app.tokenTTL")) * time.Hour)

	session := models.Session{
		UserId:     userId,
		Token:      sessionToken,
		CreatedAt:  time.Now(),
		ExpiryDate: expireDate,
	}

	if _, err := s.sessionRepository.CreateSession(session); err != nil {
		return models.Session{}, err
	}

	return session, nil
}
