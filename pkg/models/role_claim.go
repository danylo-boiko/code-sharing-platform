package models

type RoleClaimType string

const (
	OwnedRoleClaim   RoleClaimType = "owned"
	ForeignRoleClaim               = "foreign"
)

type ActionType string

const (
	ReadAction   ActionType = "read"
	CreateAction            = "create"
	UpdateAction            = "update"
	DeleteAction            = "delete"
)

type RoleClaim struct {
	Id         int    `gorm:"type:int; primaryKey; not null; autoIncrement;"`
	RoleId     int    `gorm:"type:int; not null;"`
	Role       Role   `gorm:"foreignKey:role_id;"`
	ClaimType  string `gorm:"type:varchar(30); not null;"`
	ClaimValue string `gorm:"type:varchar(20); not null;"`
}
