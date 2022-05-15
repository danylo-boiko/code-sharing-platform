package interfaces

import "code-sharing-platform/pkg/models"

type Session interface {
	GetSession(token string) (models.Session, error)
	CreateSession(session models.Session) (int, error)
}
