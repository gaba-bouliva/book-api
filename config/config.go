package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort			string
	ServerAddr			string
	DBHost					string
	DBUser 					string
	DBPort					string
	DBPass					string
	DBName					string
	DBSSLMode				string
}

var Envs = NewConfig()

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading env file")
	}
	return Config{
		ServerPort: getEnv("Port", "8081"),
		ServerAddr: getEnv("HOST", "localhost"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBUser: getEnv("Db_USER", "root"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBPass: getEnv("DB_PASS", "root"),
		DBName: getEnv("DB_NAME", "books"),
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key) 
	if !exists {
		return fallback
	}	
	return value
}