package interfaces

import (
	"code-sharing-platform/pkg/models"
	"time"
)

type Session interface {
	GetUserId(sessionToken string) (int, error)
	CreateSession(userId int) (models.Session, error)
	ExtendExpireDate(sessionToken string) (time.Time, error)
}
