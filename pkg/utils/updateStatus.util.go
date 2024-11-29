package utils

import (
	"fmt"
	"time"

	"github.com/BharaniJ27/Task-Management/internal/dto"
	"github.com/BharaniJ27/Task-Management/internal/models"
	"github.com/BharaniJ27/Task-Management/internal/services"
	"github.com/BharaniJ27/Task-Management/pkg/constants"
)

/**
 * Run the scheduled job to fetch the task and update based on criteria
 *
 */

func RunHourlyJob() {
	defer handlePanic()

	var tasks []models.Task
	tasks, error := services.GetAllTasks()

	if error != nil {
		return
	}

	for _, task := range tasks {

		go func(task models.Task) {

			// Check if it's been 1 hour since the creation
			if time.Since(task.CreatedAt) > time.Hour && task.Status != "critical" {
				var updatedTask dto.UpdateTaskDTO
				updatedTask.Description = task.Description
				updatedTask.Title = task.Title
				updatedTask.Status = constants.Status_Critical

				_, err := services.UpdateTask(updatedTask, &task)
				if err != nil {
					panic(constants.Cron_Failed)
				}
			}

		}(task)

	}
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from crashed cron job:", r)
	}
}
