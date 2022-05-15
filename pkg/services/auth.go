package services

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/repositories/interfaces"
	"golang.org/x/crypto/bcrypt"
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

func (a *AuthService) GetUser(username string) (models.User, error) {
	return a.authRepository.GetUser(username)
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
