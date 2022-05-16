package database

import (
	"code-sharing-platform/pkg/models"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"io/ioutil"
	"path/filepath"
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
	mssql, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = MigrateDatabase(mssql); err != nil {
		return nil, err
	}

	return mssql, nil
}

func MigrateDatabase(mssql *gorm.DB) error {
	mssql.AutoMigrate(models.User{}, models.Role{}, models.RoleClaim{}, models.SupportedLanguage{}, models.CodeSnippet{}, models.Session{})

	pathToSQLFiles, err := filepath.Abs("schema")
	if err != nil {
		return err
	}

	sqlFiles, err := ioutil.ReadDir(pathToSQLFiles)
	if err != nil {
		return err
	}

	for _, file := range sqlFiles {
		absPathToFile := filepath.Join(pathToSQLFiles, file.Name())
		c, ioError := ioutil.ReadFile(absPathToFile)
		if ioError != nil {
			return ioError
		}
		mssql.Exec(string(c))
	}

	return nil
}
