package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBPort    int
	JWTSecret string
}

func LoadConfig() *Config {
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}
	return &Config{
		DBPort:    port,
		JWTSecret: jwtSecret,
	}
}
