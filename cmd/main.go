package main

import (
	"github.com/BharaniJ27/Task-Management/internal/config"
	"github.com/BharaniJ27/Task-Management/pkg/cron"
	"github.com/BharaniJ27/Task-Management/pkg/routes"
	"github.com/gin-gonic/gin"
)

/**
 * Used to initialize the config before running the main application
 */
func init() {
	config.LoadEnvVariables()
	config.DatabaseInitializer()
}

/**
 * starting point of the application
 */
func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)

	go cron.CronJob()
	r.Run()

}
