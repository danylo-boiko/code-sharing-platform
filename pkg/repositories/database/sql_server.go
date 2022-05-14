package database

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

const (
	usersTable              = "users"
	usersRolesTable         = "users_roles"
	rolesTable              = "roles"
	rolesClaimsTable        = "roles_claims"
	coleSnippetsTable       = "code_snippets"
	supportedLanguagesTable = "supported_languages"
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

	return sqlServer, nil
}
