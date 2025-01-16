// main.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rohansen856/taskmanager/config"
	"github.com/rohansen856/taskmanager/routes"
)

func main() {
	config.LoadEnv()

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	r := gin.Default()
	routes.RegisterTaskRoutes(r, db)

	port := config.GetEnv("PORT", "8080")
	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}
