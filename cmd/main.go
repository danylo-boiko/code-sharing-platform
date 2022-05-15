package main

import (
	"code-sharing-platform"
	"code-sharing-platform/pkg/handlers"
	"code-sharing-platform/pkg/repositories"
	"code-sharing-platform/pkg/repositories/database"
	"code-sharing-platform/pkg/services"
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	mssql, err := database.NewSQLServer(database.Config{
		Host:         viper.GetString("mssql.host"),
		Port:         viper.GetString("mssql.port"),
		Username:     viper.GetString("mssql.username"),
		Password:     os.Getenv("SQL_SERVER_PASSWORD"),
		DatabaseName: viper.GetString("mssql.databasename"),
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}

	repositories := repositories.NewRepository(mssql)
	services := services.NewService(repositories)
	handlers := handlers.NewHandler(services)

	srv := new(code_sharing_platform.Server)
	go func() {
		if err := srv.Run(viper.GetString("app.port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Println("App started ...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("App shutting down ...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occurred on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
