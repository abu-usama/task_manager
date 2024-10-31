package database

import (
	"fmt"
	"log"
	"os"

	"task_manager/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	conf *config.Configuration
	DB   *gorm.DB
)

func ConnectDb() {
	connectToDatabase()
}

func connectToDatabase() {
	conf = config.GetEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		conf.DBHost,
		conf.DBUser,
		conf.DBPassword,
		conf.DBName,
		conf.DBPort,
		conf.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true, Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		os.Exit(2)
	}

	fmt.Println("DB connection was a success")

	DB = db
}
