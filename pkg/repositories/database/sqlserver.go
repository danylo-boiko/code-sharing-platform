package database

import (
	"code-sharing-platform/pkg/models"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

func NewSQLServer(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName)
	sqlServer, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = sqlServer.AutoMigrate(&models.User{}, &models.Role{}, &models.RoleClaim{}, &models.SupportedLanguage{}, &models.CodeSnippet{}, &models.RefreshToken{})

	if err != nil {
		return nil, err
	}

	return sqlServer, nil
}
