package cron

import (
	"time"

	"github.com/BharaniJ27/Task-Management/pkg/utils"
)

/**
 * Run the scheduled job every 10 minutes
 *
 */

func CronJob() {

	ticker := time.NewTicker(10 * time.Minute) // Trigger every 10 Minutes
	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				utils.RunHourlyJob()
			}
		}
	}()
}
