package config

import (
	"fmt"
	"os"
)

type Config struct {
	PORT    string
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PSW  string
	DB_NAME string
}

func Load() *Config {
	return &Config{
		PORT:    getEnv("PORT", "8080"),
		DB_HOST: getEnv("DB_HOST", "localhost"),
		DB_PORT: getEnv("DB_PORT", "3306"),
		DB_USER: getEnv("DB_USER", "root"),
		DB_PSW:  getEnv("DB_PSW", "Zjk007326"),
		DB_NAME: getEnv("DB_NAME", "schoolmarket"),
	}
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DB_USER, c.DB_PSW, c.DB_HOST, c.DB_PORT, c.DB_NAME,
	)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
