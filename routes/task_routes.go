// routes/task_routes.go
package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rohansen856/taskmanager/handlers"
)

// RegisterTaskRoutes sets up the task routes.
func RegisterTaskRoutes(router *gin.Engine, db *sql.DB) {
	taskGroup := router.Group("/tasks")
	{
		taskGroup.POST("/", handlers.CreateTaskHandler(db))
		taskGroup.GET("/", handlers.GetTasksHandler(db))
		taskGroup.GET("/:id", handlers.GetTaskHandler(db))
		taskGroup.PATCH("/:id", handlers.UpdateTaskHandler(db))
		taskGroup.DELETE("/:id", handlers.DeleteTaskHandler(db))
	}
}
