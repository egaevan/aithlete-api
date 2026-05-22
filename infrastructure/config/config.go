package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
}

type AuthConfig struct {
	JWTSecret     string
	JWTExpiration int
}

type ServerConfig struct {
	Host     string
	Port     int
	Env      string
	TLS      TLSConfig
}

type TLSConfig struct {
	Enabled  bool
	CertFile string
	KeyFile  string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnvAsInt("SERVER_PORT", 8080),
			Env:  getEnv("APP_ENV", "development"),
			TLS: TLSConfig{
				Enabled:  getEnv("TLS_ENABLED", "false") == "true",
				CertFile: getEnv("TLS_CERT_FILE", "./cert.pem"),
				KeyFile:  getEnv("TLS_KEY_FILE", "./key.pem"),
			},
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "aithlete"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Auth: AuthConfig{
			JWTSecret:     getEnv("JWT_SECRET", "dev-secret"),
			JWTExpiration: getEnvAsInt("JWT_EXPIRATION_HOURS", 24),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valStr := getEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}
