package interfaces

import (
	"code-sharing-platform/pkg/models"
)

type Authorization interface {
	GetUserById(id int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	CreateUser(user models.User) (int, error)
	HashPassword(password string) string
	IsPasswordCorrect(password, hash string) bool
}
