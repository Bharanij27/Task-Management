package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

/**
 * Func: LoadEnvVariables is to load the env from the .env file
 */

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

/**
 * Func: GetEnv is to fetch the env value stored in .env
 *
 * @params key: key name
 * @params defaultValue: default value incase the given key doesn't have value in .env
 * @return value for the given key
 */

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
