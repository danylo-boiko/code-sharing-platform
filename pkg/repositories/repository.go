package repositories

import (
	"code-sharing-platform/pkg/models"
	"gorm.io/gorm"
)

type Authorization interface {
	SignIn(username string, password string) (models.User, error)
	SignUp(user *models.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(sqlServer *gorm.DB) *Repository {
	return &Repository{}
}
