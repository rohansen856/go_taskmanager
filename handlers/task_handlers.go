package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func CreateTaskHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		query := "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id"
		err := db.QueryRow(query, task.Title, task.Description, task.Status).Scan(&task.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, task)
	}
}

func GetTasksHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, description, status FROM tasks")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		tasks := []Task{}
		for rows.Next() {
			var task Task
			if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			tasks = append(tasks, task)
		}
		c.JSON(http.StatusOK, tasks)
	}
}

func GetTaskHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var task Task
		query := "SELECT id, title, description, status FROM tasks WHERE id = $1"
		if err := db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, task)
	}
}

func UpdateTaskHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var task Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		query := "UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4"
		result, err := db.Exec(query, task.Title, task.Description, task.Status, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		numRows, _ := result.RowsAffected()
		if numRows == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
	}
}

func DeleteTaskHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		query := "DELETE FROM tasks WHERE id = $1"
		result, err := db.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		numRows, _ := result.RowsAffected()
		if numRows == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
	}
}
