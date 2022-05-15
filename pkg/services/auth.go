package services

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/repositories/interfaces"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	authRepository    interfaces.Authorization
	sessionRepository interfaces.Session
}

func NewAuthService(authRepository interfaces.Authorization, sessionRepository interfaces.Session) *AuthService {
	return &AuthService{
		authRepository:    authRepository,
		sessionRepository: sessionRepository,
	}
}

func (a *AuthService) CreateSessionToken(username, password string) (string, time.Time, error) {
	user, err := a.authRepository.GetUser(username)
	if err != nil {
		return "", time.Time{}, err
	}

	if isPasswordCorrect := a.IsPasswordCorrect(password, user.PasswordHash); !isPasswordCorrect {
		return "", time.Time{}, errors.New("provided wrong password")
	}

	sessionToken := uuid.NewString()
	expireDate := time.Now().UTC().Add(1 * time.Hour)

	session := models.Session{
		UserId:     user.Id,
		Token:      sessionToken,
		CreatedAt:  time.Now(),
		ExpiryDate: expireDate,
	}

	if _, err := a.sessionRepository.CreateSession(session); err != nil {
		return "", time.Time{}, errors.New("session creating error")
	}

	return sessionToken, expireDate, nil
}

func (a *AuthService) CreateUser(user models.User) (int, error) {
	return a.authRepository.CreateUser(user)
}

func (a *AuthService) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (a *AuthService) IsPasswordCorrect(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
