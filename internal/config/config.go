package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

var instance *Config
var once sync.Once

func LoadEnv() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using default environment values.")
	}

	once.Do(func() {
		instance = &Config{
			Port: GetEnv("PORT", "8080"),
		}
	})
}

func GetConfig() *Config {
	return instance
}

func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
