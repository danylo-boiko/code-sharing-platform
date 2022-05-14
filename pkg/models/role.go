package models

type Role struct {
	Id          int         `gorm:"type:int; primaryKey; not null; autoIncrement;"`
	Title       string      `gorm:"type:varchar(50); not null; unique;"`
	Claims      []RoleClaim `gorm:"foreignKey:role_id;"`
	Description string      `gorm:"type:varchar(100);"`
}
