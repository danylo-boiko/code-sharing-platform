package repositories

import (
	"gorm.io/gorm"
)

type AuthSqlServer struct {
	sqlServer *gorm.DB
}

func NewAuthSqlServer(sqlServer *gorm.DB) *AuthSqlServer {
	return &AuthSqlServer{sqlServer: sqlServer}
}
