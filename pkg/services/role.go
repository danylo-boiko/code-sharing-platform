package services

import (
	"code-sharing-platform/pkg/models"
	"code-sharing-platform/pkg/repositories/interfaces"
	"errors"
)

type RoleService struct {
	roleRepository interfaces.Role
}

func NewRoleService(roleRepository interfaces.Role) *RoleService {
	return &RoleService{roleRepository: roleRepository}
}

func (r *RoleService) GetUserRoles(userId int) ([]models.Role, error) {
	return r.roleRepository.GetUserRoles(userId)
}

func (r *RoleService) CheckUserPermission(userId int, roleClaim models.RoleClaimType, action models.ActionType) error {
	roles, err := r.roleRepository.GetUserRoles(userId)
	if err != nil {
		return err
	}

	for _, role := range roles {
		for _, claim := range role.Claims {
			if claim.ClaimType == string(roleClaim) && claim.ClaimValue == string(action) {
				return nil
			}
		}
	}

	return errors.New("you don't have permission for this operation")
}
