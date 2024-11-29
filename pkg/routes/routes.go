package routes

import (
	"github.com/BharaniJ27/Task-Management/api/controllers"
	"github.com/gin-gonic/gin"
)

/**
 * Func: RegisterRoutes is for registering the routes
 *
 * @params c gin.Enginer - The Gin Enginer to group and add API
 */

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/tasks", controllers.GetTasks)
		api.GET("/tasks/:id", controllers.GetTaskById)
		api.POST("/task", controllers.CreateTask)
		api.PUT("/tasks/:id", controllers.UpdateTask)
		api.DELETE("/tasks/:id", controllers.DeleteTask)
	}
}
