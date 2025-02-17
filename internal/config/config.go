package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	LogLevel   string
	AppName    string
	AppHost    string
	AppPort    string
	AppEnv     string
	DbType     string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
	DbFilepath string
	DbSslMode  string
}

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("No .env file found, falling back to environment variables")
	}

	config := &Config{
		LogLevel: getEnv("LOG_LEVEL", zerolog.DebugLevel.String()),
		AppName:  getEnv("APP_NAME", "Product list"),
		AppHost:  getEnv("APP_HOST", "localhost"),
		AppPort:  getEnv("APP_PORT", "8080"),
		AppEnv:   getEnv("APP_ENV", "local"),

		//DB config
		DbType:     getEnv("DB_TYPE", "mysql"),
		DbUser:     getEnv("DB_USER", "root"),
		DbPassword: getEnv("DB_PASSWORD", ""),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "3306"),
		DbName:     getEnv("DB_NAME", "product_list"),
		DbFilepath: getEnv("DB_FILE_PATH", "./.temp/sqlite/product_list.db"), //sqlite
		DbSslMode:  getEnv("DB_SSL_MODE", "disable"),                         //postgres
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
