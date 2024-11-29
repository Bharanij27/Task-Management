package main

import (
	"github.com/BharaniJ27/Task-Management/internal/config"
	"github.com/BharaniJ27/Task-Management/internal/models"
)

/**
 * Used to initialize the config before running the main application
 */
func init() {
	config.LoadEnvVariables()
	config.DatabaseInitializer()
}

/**
 * main function of package will do migration for given models
 */

func main() {
	config.DB.AutoMigrate(&models.Task{})
}
