package interfaces

import "code-sharing-platform/pkg/models"

type Session interface {
	GetUserId(sessionToken string) (int, error)
	CreateSession(userId int) (models.Session, error)
}
