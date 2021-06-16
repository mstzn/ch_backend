package config

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnvironmentVariable(key string, def string) string  {
	err := godotenv.Load()
	if err != nil {
		panic("Can not load environment file")
	}
	envValue := os.Getenv(key)

	if envValue == "" {
		return def
	}

	return envValue
}
