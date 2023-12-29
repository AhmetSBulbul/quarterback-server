package helpers

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvOrDefault(key string, def string) string {
	value := os.Getenv(key)
	if len(value) > 0 {
		return value
	}
	return def
}

func GetEnvOrFail(key string) (string, error) {
	value := os.Getenv(key)
	if len(value) > 0 {
		return value, nil
	}
	return "", errors.New("cannot read " + key)
}

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	return nil
}
