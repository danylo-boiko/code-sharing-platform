package interfaces

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/requests/auth"
)

type Authorization interface {
	GetUserById(id int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	CreateUser(request auth.SignUpRequest) (int, error)
	HashPassword(password string) string
	IsPasswordCorrect(password, hash string) bool
}
