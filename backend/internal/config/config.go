package config

import "os"

type Config struct {
    Port        string
    DatabaseURL string
    JWTSecret   string
    RedisURL    string
}

func Load() *Config {
    return &Config{
        Port:        getEnv("PORT", "8080"),
        DatabaseURL: getEnv("DATABASE_URL", "postgres://erp:erp123@localhost:5432/school_erp?sslmode=disable"),
        JWTSecret:   getEnv("JWT_SECRET", "secret"),
        RedisURL:    getEnv("REDIS_URL", "localhost:6379"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}