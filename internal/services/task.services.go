package services

import (
	"github.com/BharaniJ27/Task-Management/internal/config"
	"github.com/BharaniJ27/Task-Management/internal/dto"
	"github.com/BharaniJ27/Task-Management/internal/models"
	"gorm.io/gorm"
)

/**
 * Func: GetAllTasks is used to fetch all task from DB
 *
 * @return tasks list and error if any
 */

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task

	// Using a transaction to fetch tasks
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Find(&tasks)
		return result.Error
	})

	return tasks, err
}

/**
 * Func: GetTaskById is used to get a single task from DB using Id
 *
 * @param taskId: taskId of the requested task
 * @return tasks list and error if any
 */

func GetTaskById(taskId string) (models.Task, error) {
	var tasks models.Task
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.First(&tasks, taskId)
		return result.Error
	})
	return tasks, err
}

/**
 * Func: CreateTask is used to fetch all task from DB
 * @param body - Request body which contains tasks
 *
 * @return tasks list and error if any
 */

func CreateTask(body dto.CreateTaskDTO) (models.Task, error) {

	task := models.Task{
		Title:       body.Title,
		Description: body.Description,
		Status:      body.Status,
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Save the task to the database
		result := tx.Create(&task)
		return result.Error
	})

	// Respond with the created task
	return task, err
}

/**
 * Func: UpdateTask is used update a single task
 *
 * @param body - Request body which contains tasks
 * @param task - task that needs to be updated
 * @return tasks list and error if any
 */

func UpdateTask(body dto.UpdateTaskDTO, task *models.Task) (models.Task, error) {

	updatedTask := models.Task{
		Title:       body.Title,
		Description: body.Description,
		Status:      body.Status,
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Update the task to the database
		result := tx.Model(task).Updates(updatedTask)
		return result.Error
	})

	// Respond with the created task
	return updatedTask, err
}

/**
 * Func: DeleteTask is used to soft delete a single task from DB
 * @param taskId: taskId of the requested task
 *
 * @return tasks list and error if any
 */

func DeleteTask(taskId string) error {

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Save the task to the database
		result := tx.Delete(&models.Task{}, taskId)
		return result.Error
	})
	return err
}
