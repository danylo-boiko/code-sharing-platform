package models

import "time"

type Session struct {
	Id         int       `gorm:"type:int; primaryKey; not null; autoIncrement;"`
	UserId     int       `gorm:"type:int; not null;"`
	User       User      `gorm:"foreignKey:user_id;"`
	Token      string    `gorm:"type:varchar(36); not null; unique;"`
	CreatedAt  time.Time `gorm:"type:datetime2; not null;"`
	ExpiryDate time.Time `gorm:"type:datetime2; not null;"`
}
