package interfaces

import (
	"code-sharing-platform/pkg/models"
	"time"
)

type Authorization interface {
	CreateSessionToken(username, password string) (string, time.Time, error)
	CreateUser(user models.User) (int, error)
	HashPassword(password string) string
	IsPasswordCorrect(password, hash string) bool
}
