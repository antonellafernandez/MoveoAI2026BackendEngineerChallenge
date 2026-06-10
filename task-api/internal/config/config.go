package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
	Auth     AuthConfig
	Frontend FrontendConfig
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

type FrontendConfig struct {
	URL string
}

func Load() *Config {
	_ = godotenv.Load()

	dbPort := getEnvInt("DB_PORT", 5432)
	serverPort := getEnvInt("SERVER_PORT", 8080)

	cfg := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "tasks"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
		Server: ServerConfig{
			Port: serverPort,
		},
		Auth: AuthConfig{
			AdminUsername: os.Getenv("ADMIN_USERNAME"),
			AdminPassword: os.Getenv("ADMIN_PASSWORD"),
		},
		Frontend: FrontendConfig{
			URL: getEnv("FRONTEND_URL", "http://localhost:5173"),
		},
	}

	validateConfig(cfg)

	return cfg
}

func validateConfig(cfg *Config) {
	if cfg.JWT.Secret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	if cfg.Auth.AdminUsername == "" {
		log.Fatal("ADMIN_USERNAME is required")
	}

	if cfg.Auth.AdminPassword == "" {
		log.Fatal("ADMIN_PASSWORD is required")
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
