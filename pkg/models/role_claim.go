package models

type RoleClaim struct {
	Id         int    `gorm:"type:int; primaryKey; not null; autoIncrement;"`
	RoleId     int    `gorm:"type:int; not null;"`
	Role       Role   `gorm:"foreignKey:role_id;"`
	ClaimType  string `gorm:"type:varchar(30); not null;"`
	ClaimValue string `gorm:"type:varchar(20); not null;"`
}
