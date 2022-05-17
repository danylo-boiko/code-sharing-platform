package services

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/repositories/interfaces"
	"code-sharing-platform/pkg/requests/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository interfaces.Authorization
	roleRepository interfaces.Role
}

func NewAuthService(authRepository interfaces.Authorization, roleRepository interfaces.Role) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		roleRepository: roleRepository,
	}
}

func (a *AuthService) GetUserById(id int) (models.User, error) {
	return a.authRepository.GetUserById(id)
}

func (a *AuthService) GetUserByUsername(username string) (models.User, error) {
	return a.authRepository.GetUserByUsername(username)
}

func (a *AuthService) CreateUser(request auth.SignUpRequest) (int, error) {
	defaultRole, err := a.roleRepository.GetRole(models.DefaultUserRole)
	if err != nil {
		return 0, err
	}

	userId, err := a.authRepository.CreateUser(models.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: a.HashPassword(request.Password),
		Roles:        []models.Role{defaultRole},
	})
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (a *AuthService) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (a *AuthService) IsPasswordCorrect(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
