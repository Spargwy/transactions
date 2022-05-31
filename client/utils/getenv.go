package utils

import (
	"log"
	"os"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("%s not set, used default value: %s", key, defaultValue)
		return defaultValue
	}

	return value
}
