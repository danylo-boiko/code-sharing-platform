package repositories

import (
	"code-sharing-platform/pkg/models"
	"gorm.io/gorm"
)

type AuthMsSQL struct {
	mssql *gorm.DB
}

func NewAuthMsSQL(mssql *gorm.DB) *AuthMsSQL {
	return &AuthMsSQL{mssql: mssql}
}

func (a *AuthMsSQL) GetUserById(id int) (models.User, error) {
	var user models.User
	if err := a.mssql.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (a *AuthMsSQL) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := a.mssql.First(&user, "username = ?", username).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (a *AuthMsSQL) CreateUser(user models.User) (int, error) {
	if err := a.mssql.Create(&user).Error; err != nil {
		return user.Id, err
	}
	return user.Id, nil
}
