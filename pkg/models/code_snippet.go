package models

import "time"

type CodeSnippet struct {
	Id         int               `gorm:"type:int; primaryKey; not null; autoIncrement;"`
	UserId     int               `gorm:"type:int; not null;"`
	User       User              `gorm:"foreignKey:user_id;"`
	LanguageId int               `gorm:"type:int; not null;"`
	Language   SupportedLanguage `gorm:"foreignKey:language_id;"`
	Header     string            `gorm:"type:varchar(50); not null;"`
	Code       string            `gorm:"type:varchar(MAX); not null;"`
	CreatedAt  time.Time         `gorm:"type:datetime2; not null;"`
	ExpiryDate time.Time         `gorm:"type:datetime2;"`
	ViewsLimit int               `gorm:"type:int; check:views_limit>=0;"`
	Views      int               `gorm:"type:int; not null; default:0; check:views>=0;"`
}
