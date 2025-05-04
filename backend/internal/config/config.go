package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds the application configuration
type Config struct {
	Env      string
	Server   ServerConfig
	Database DatabaseConfig
	CORS     CORSConfig
	Svea     SveaConfig
	RateLimit RateLimitConfig
}

// ServerConfig holds the HTTP server configuration
type ServerConfig struct {
	Port                  int
	ReadTimeoutSeconds    int
	WriteTimeoutSeconds   int
	IdleTimeoutSeconds    int
	ShutdownTimeoutSeconds int
}

// DatabaseConfig holds the database configuration
type DatabaseConfig struct {
	Driver        string
	DSN           string
	MaxOpenConns  int
	MaxIdleConns  int
	ConnMaxLifetime time.Duration
	RunMigrations bool
	MigrationsPath string
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
	MaxAge         int
}

// SveaConfig holds Svea Ekonomi API configuration
type SveaConfig struct {
	BaseURL     string
	MerchantID  string
	Secret      string
	Environment string
	Timeout     time.Duration
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	RequestsPerMinute int
	BurstSize         int
	Enabled           bool
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	config := &Config{
		Env: getEnv("APP_ENV", "development"),
		Server: ServerConfig{
			Port:                   getEnvAsInt("SERVER_PORT", 8080),
			ReadTimeoutSeconds:     getEnvAsInt("SERVER_READ_TIMEOUT", 30),
			WriteTimeoutSeconds:    getEnvAsInt("SERVER_WRITE_TIMEOUT", 30),
			IdleTimeoutSeconds:     getEnvAsInt("SERVER_IDLE_TIMEOUT", 60),
			ShutdownTimeoutSeconds: getEnvAsInt("SERVER_SHUTDOWN_TIMEOUT", 30),
		},
		Database: DatabaseConfig{
			Driver:          getEnv("DB_DRIVER", "mysql"),
			DSN:             buildDSN(),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 25),
			ConnMaxLifetime: time.Duration(getEnvAsInt("DB_CONN_MAX_LIFETIME", 5)) * time.Minute,
			RunMigrations:   getEnvAsBool("DB_RUN_MIGRATIONS", true),
			MigrationsPath:  getEnv("DB_MIGRATIONS_PATH", "migrations"),
		},
		CORS: CORSConfig{
			AllowedOrigins: getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"*"}),
			AllowedMethods: getEnvAsSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			AllowedHeaders: getEnvAsSlice("CORS_ALLOWED_HEADERS", []string{"Origin", "Content-Type", "Accept", "Authorization"}),
			MaxAge:         getEnvAsInt("CORS_MAX_AGE", 86400),
		},
		Svea: SveaConfig{
			BaseURL:     getEnv("SVEA_BASE_URL", "https://checkoutapistage.svea.com"),
			MerchantID:  getEnv("SVEA_MERCHANT_ID", ""),
			Secret:      getEnv("SVEA_SECRET", ""),
			Environment: getEnv("SVEA_ENVIRONMENT", "stage"),
			Timeout:     time.Duration(getEnvAsInt("SVEA_TIMEOUT", 10)) * time.Second,
		},
		RateLimit: RateLimitConfig{
			RequestsPerMinute: getEnvAsInt("RATE_LIMIT_REQUESTS_PER_MINUTE", 60),
			BurstSize:         getEnvAsInt("RATE_LIMIT_BURST_SIZE", 10),
			Enabled:           getEnvAsBool("RATE_LIMIT_ENABLED", true),
		},
	}

	// Validate required configuration
	if config.Svea.MerchantID == "" {
		return nil, fmt.Errorf("SVEA_MERCHANT_ID is required")
	}

	if config.Svea.Secret == "" {
		return nil, fmt.Errorf("SVEA_SECRET is required")
	}

	return config, nil
}

// Helper function to build database DSN from component environment variables
func buildDSN() string {
	driver := getEnv("DB_DRIVER", "mysql")
	
	if dsn := getEnv("DB_DSN", ""); dsn != "" {
		return dsn
	}

	switch driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			getEnv("DB_USER", "root"),
			getEnv("DB_PASSWORD", ""),
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "3306"),
			getEnv("DB_NAME", "svenskhalsovard"),
		)
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "5432"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", ""),
			getEnv("DB_NAME", "svenskhalsovard"),
		)
	default:
		return ""
	}
}

// Helper functions to read environment variables

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.ParseBool(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	if valueStr, exists := os.LookupEnv(key); exists && valueStr != "" {
		return strings.Split(valueStr, ",")
	}
	return defaultValue
}