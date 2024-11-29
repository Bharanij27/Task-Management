package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BharaniJ27/Task-Management/internal/dto"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.Internal_Error})
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
		message := fmt.Sprintf(constants.Invalid_Task_Fetch, taskId)
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.Internal_Error})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.Invalid_Payload})
		return
	}

	task, err := services.CreateTask(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.Internal_Error})
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
		message := fmt.Sprintf(constants.Invalid_Task_Update, taskId)
		c.JSON(http.StatusForbidden, gin.H{"error": message})
		return
	}

	// get the task detail from requets payload
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.Invalid_Payload})
		return
	}

	_, onUpdateErr := services.UpdateTask(body, &task)
	if onUpdateErr != nil {
		fmt.Println(onUpdateErr.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.Internal_Error})
		return
	}
	message := fmt.Sprintf(constants.Update_Success, taskId)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("id")

	err := services.DeleteTask(taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.Deletion_Failed})
		return
	}
	message := fmt.Sprintf(constants.Deletion_Success, taskId)
	c.JSON(http.StatusOK, gin.H{"message": message})
}
