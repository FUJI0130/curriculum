// src/core/config/env.go

package config

import "os"

type EnvConfig struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
	AppPort    string
}

func LoadEnv() *EnvConfig {
	return &EnvConfig{
		DbUser:     getEnv("DB_USER", "user"),
		DbPassword: getEnv("DB_PASSWORD", "password"),
		DbHost:     getEnv("DB_HOST", "mysql"),
		DbPort:     getEnv("DB_PORT", "3306"),
		DbName:     getEnv("DB_NAME", "sql"),
		AppPort:    getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
