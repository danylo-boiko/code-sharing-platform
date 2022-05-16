package interfaces

import (
	"code-sharing-platform/pkg/models"
	"time"
)

type Session interface {
	GetUserId(sessionToken string) (int, error)
	GetSession(sessionToken string) (models.Session, error)
	CreateSession(userId int) (models.Session, error)
	ExtendExpireDate(sessionToken string) (time.Time, error)
}
