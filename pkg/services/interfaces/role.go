package interfaces

import "code-sharing-platform/pkg/models"

type Role interface {
	GetUserRoles(userId int) ([]models.Role, error)
	CheckUserPermission(userId int, roleClaim models.RoleClaimType, action models.ActionType) error
}
