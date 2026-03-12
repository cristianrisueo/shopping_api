package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config aggregates all application configuration sections.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	AWS      AWSConfig
	Upload   UploadConfig
}

// ServerConfig holds HTTP server settings.
type ServerConfig struct {
	Port    string
	GinMode string
}

// DatabaseConfig holds PostgreSQL connection settings.
type DatabaseConfig struct {
	DSN      string
	Port     string
	Host     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// JWTConfig holds token signing and expiration settings.
type JWTConfig struct {
	SecretKey              string
	Expiration             time.Duration
	RefreshTokenExpiration time.Duration
}

// AWSConfig holds AWS credentials and S3 settings.
type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	S3Endpoint      string
	S3BucketName    string
}

// UploadConfig holds file upload path and size limit settings.
type UploadConfig struct {
	Path    string
	MaxSize int64
}

// LoadConfig loads the .env file and returns a populated Config.
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("SERVER_PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "release"),
		},
		Database: DatabaseConfig{
			DSN:      getEnv("DB_DSN", ""),
			Port:     getEnv("DB_PORT", "5432"),
			Host:     getEnv("DB_HOST", "localhost"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "shopping"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			SecretKey:              getEnv("JWT_SECRET_KEY", "secret-key"),
			Expiration:             time.Duration(getEnvInt("JWT_EXPIRATION", 3600)) * time.Second,
			RefreshTokenExpiration: time.Duration(getEnvInt("JWT_REFRESH_TOKEN_EXPIRATION", 86400)) * time.Second,
		},
		AWS: AWSConfig{
			Region:          getEnv("AWS_REGION", "us-east-1"),
			AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", ""),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
			S3Endpoint:      getEnv("S3_ENDPOINT", ""),
			S3BucketName:    getEnv("S3_BUCKET_NAME", ""),
		},
		Upload: UploadConfig{
			Path:    getEnv("UPLOAD_PATH", "./uploads"),
			MaxSize: int64(getEnvInt("UPLOAD_MAX_SIZE", 1048576)),
		},
	}, nil
}

// getEnv returns the environment variable value or a default if not set.
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

// getEnvInt returns the environment variable parsed as int, or a default if not set or invalid.
func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}
