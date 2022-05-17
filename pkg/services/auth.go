package services

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/repositories/interfaces"
	"code-sharing-platform/pkg/requests/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository interfaces.Authorization
}

func NewAuthService(authRepository interfaces.Authorization) *AuthService {
	return &AuthService{authRepository: authRepository}
}

func (a *AuthService) GetUserById(id int) (models.User, error) {
	return a.authRepository.GetUserById(id)
}

func (a *AuthService) GetUserByUsername(username string) (models.User, error) {
	return a.authRepository.GetUserByUsername(username)
}

func (a *AuthService) CreateUser(request auth.SignUpRequest) (int, error) {
	return a.authRepository.CreateUser(models.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: a.HashPassword(request.Password),
	})
}

func (a *AuthService) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (a *AuthService) IsPasswordCorrect(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
