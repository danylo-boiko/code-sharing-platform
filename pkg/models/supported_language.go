package models

type SupportedLanguage struct {
	Id            int    `gorm:"type:int; primaryKey; not null; autoIncrement;"`
	Title         string `gorm:"type:varchar(20); not null; unique;"`
	FileExtension string `gorm:"type:varchar(10); not null; unique;"`
}
