package repositories

import (
	"code-sharing-platform/pkg/models"
	"gorm.io/gorm"
)

type RoleMsSQL struct {
	mssql *gorm.DB
}

func NewRoleMsSQL(mssql *gorm.DB) *RoleMsSQL {
	return &RoleMsSQL{mssql: mssql}
}

func (r *RoleMsSQL) GetUserRoles(userId int) ([]models.Role, error) {
	var rolesIds []int
	err := r.mssql.Table("users_roles").Select("role_id").Where("user_id = ?", userId).Scan(&rolesIds).Error
	if err != nil {
		return nil, err
	}

	var userRoles []models.Role
	if err := r.mssql.Preload("Claims").Where(rolesIds).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}
