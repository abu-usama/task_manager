package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// Environtments: Development, Staging, Production

type Configuration struct {
	Host       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
	Port       string
	// SwaggerHost string
}

var (
	once           sync.Once
	configInstance *Configuration
)

func GetEnv() *Configuration {
	once.Do(func() {
		env := os.Getenv("APP_ENV")

		fmt.Println("Enviroment Loded: " + env)

		if env == "" {
			env = "development"
		}

		loc := "./config/.env." + env
		err := godotenv.Load(loc)
		if err != nil {
			log.Fatalf("Error loading ./config/.env.%s file: %v", env, err)
		}

		configInstance = &Configuration{
			DBHost:     os.Getenv("DB_HOST"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			DBPort:     os.Getenv("DB_PORT"),
			DBSSLMode:  os.Getenv("DB_SSLMODE"),
			Port:       os.Getenv("PORT"),
			// SwaggerHost: os.Getenv("SWAGGER_HOST"),
		}

		// Override SwaggerHost for development environment
		// if env == "development" {
		// 	configInstance.SwaggerHost = "http://localhost:" + configInstance.Port
		// }
	})
	return configInstance
}