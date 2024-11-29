package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BharaniJ27/Task-Management/internal/dto"
	errorHandler "github.com/BharaniJ27/Task-Management/internal/error"
	"github.com/BharaniJ27/Task-Management/internal/services"
	"github.com/BharaniJ27/Task-Management/pkg/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
 * Func: GetTask is a controller function to fetch the all task details
 *
 * @params c gin.Context - The Gin context for handling the HTTP request and response.
 * @return null
 */

func GetTasks(c *gin.Context) {
	tasks, err := services.GetAllTasks()
	if err != nil {
		errorHandler.ErrorHandler(c, constants.Internal_Error, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

/**
 * Func: GetTaskById is a controller function to fetch a single task details
 *
 * @params c gin.Context - The Gin context for handling the HTTP request and response.
 * @return null
 */

func GetTaskById(c *gin.Context) {
	taskId := c.Param("id")
	task, err := services.GetTaskById(taskId)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		errorHandler.ErrorHandler(c, constants.Invalid_Task_Fetch, http.StatusBadRequest)
		return
	}

	if err != nil {
		errorHandler.ErrorHandler(c, constants.Internal_Error, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, task)
}

/**
 * Func: CreateTask is a controller function to add a new task
 *
 * @params c gin.Context - The Gin context for handling the HTTP request and response.
 * @return null
 */

func CreateTask(c *gin.Context) {
	var body dto.CreateTaskDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		errorHandler.ErrorHandler(c, constants.Invalid_Payload, http.StatusBadRequest)
		return
	}

	task, err := services.CreateTask(body)
	if err != nil {
		errorHandler.ErrorHandler(c, constants.Internal_Error, http.StatusInternalServerError)
		return
	}
	message := fmt.Sprintf(constants.Creation_Success, strconv.Itoa(int(task.ID)))
	c.JSON(http.StatusOK, gin.H{"message": message})
}

/**
 * Func: UpdateTask is a controller function to update a task
 *
 * @params c gin.Context - The Gin context for handling the HTTP request and response.
 * @return null
 */

func UpdateTask(c *gin.Context) {
	var body dto.UpdateTaskDTO
	taskId := c.Param("id")

	// Check if the task exists
	task, err := services.GetTaskById(taskId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errorHandler.ErrorHandler(c, constants.Invalid_Task_Update, http.StatusForbidden)
		return
	}

	// get the task detail from requets payload
	if err := c.ShouldBindJSON(&body); err != nil {
		errorHandler.ErrorHandler(c, constants.Invalid_Payload, http.StatusBadRequest)
		return
	}

	_, onUpdateErr := services.UpdateTask(body, &task)
	if onUpdateErr != nil {
		errorHandler.ErrorHandler(c, constants.Internal_Error, http.StatusInternalServerError)
		return
	}
	message := fmt.Sprintf(constants.Update_Success, taskId)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("id")

	err := services.DeleteTask(taskId)
	if err != nil {
		errorHandler.ErrorHandler(c, constants.Deletion_Failed, http.StatusInternalServerError)
		return
	}
	message := fmt.Sprintf(constants.Deletion_Success, taskId)
	c.JSON(http.StatusOK, gin.H{"message": message})
}
