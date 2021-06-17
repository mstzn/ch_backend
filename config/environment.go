package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func GetEnvironmentVariable(key string, def string) string {
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Can not load environment file")
	}
	envValue := os.Getenv(key)

	if envValue == "" {
		return def
	}

	return envValue
}
