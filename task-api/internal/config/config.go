package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
	Auth     AuthConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret string
}

type ServerConfig struct {
	Port int
}

type AuthConfig struct {
	AdminUsername string
	AdminPassword string
}

func Load() *Config {
	_ = godotenv.Load()

	dbPort := getEnvInt("DB_PORT", 5432)
	serverPort := getEnvInt("SERVER_PORT", 8080)

	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "tasks"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "my-secret-key"),
		},
		Server: ServerConfig{
			Port: serverPort,
		},
		Auth: AuthConfig{
			AdminUsername: getEnv("ADMIN_USERNAME", "admin"),
			AdminPassword: getEnv("ADMIN_PASSWORD", "admin"),
		},
	}
}

func (d *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host,
		d.Port,
		d.User,
		d.Password,
		d.DBName,
		d.SSLMode,
	)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}
