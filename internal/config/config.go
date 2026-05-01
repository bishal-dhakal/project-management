package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_Host     string
    DB_Port     string
    DB_User     string
    DB_Password string
    DB_Name     string
    DB_SSLMode  string
	Server_Port string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: No .env file found, using system envs.")
	}

	log.Println(".env file found, using system envs.")

	return &Config{
		DB_Host:  getEnv("DB_HOST",""),
		DB_Port: getEnv("DB_PORT",""),
		DB_User:  getEnv("DB_USER",""),
		DB_Password: getEnv("DB_PASSWORD",""),
		DB_Name:  getEnv("DB_NAME",""),
		Server_Port: getEnv("SERVER_PORT",""),
	}
}

func getEnv(key, fallback string) string{
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}