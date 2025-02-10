package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbHost := viper.GetString("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbName := viper.GetString("DB_NAME")
	if dbName == "" {
		dbName = "followers"
	}
	dbUser := viper.GetString("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := viper.GetString("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		dbHost, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
