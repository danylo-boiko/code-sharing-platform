package models

import "time"

type RefreshToken struct {
	Id         int       `gorm:"type:int; primaryKey; not null; autoIncrement;"`
	UserId     int       `gorm:"type:int; not null;"`
	User       User      `gorm:"foreignKey:user_id;"`
	CreatedAt  time.Time `gorm:"type:datetime2; not null;"`
	ExpiryDate time.Time `gorm:"type:datetime2; not null;"`
}
