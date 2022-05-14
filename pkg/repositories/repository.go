package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
}

func NewRepository(sqlServer *gorm.DB) *Repository {
	return &Repository{}
}
