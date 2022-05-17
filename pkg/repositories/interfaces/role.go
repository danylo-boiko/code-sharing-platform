package interfaces

import "code-sharing-platform/pkg/models"

type Role interface {
	GetUserRoles(userId int) ([]models.Role, error)
}
