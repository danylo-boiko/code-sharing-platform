package repositories

import (
	"code-sharing-platform/pkg/models"
	"gorm.io/gorm"
)

type AuthMsSql struct {
	mssql *gorm.DB
}

func NewAuthMsSQL(mssql *gorm.DB) *AuthMsSql {
	return &AuthMsSql{mssql: mssql}
}

func (a *AuthMsSql) GetUser(username string) (models.User, error) {
	var user models.User
	if err := a.mssql.First(&user, "username = ?", username).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (a *AuthMsSql) CreateUser(user models.User) (int, error) {
	if err := a.mssql.Create(&user).Error; err != nil {
		return user.Id, err
	}
	return user.Id, nil
}
