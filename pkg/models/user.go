package models

type User struct {
	Id           int    `gorm:"type:int; primaryKey; not null; autoIncrement;"`
	Username     string `gorm:"type:varchar(20); not null; unique;"`
	Email        string `gorm:"type:varchar(50); not null; unique;"`
	PasswordHash string `gorm:"type:varchar(150); not null;"`
	Roles        []Role `gorm:"many2many:users_roles;"`
}
