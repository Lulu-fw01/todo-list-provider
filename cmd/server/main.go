package main

import (
	"log"
	"net/http"
	"runtime/internal/sys"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Add middleware for CORS and logging
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// API routes
	api := r.Group("/api/v1")
	{
		// Todo routes
		todos := api.Group("/todos")
		{
			todos.GET("/", getTodos)
			todos.GET("/:id", getTodo)
			todos.POST("/", createTodo)
			todos.PUT("/:id", updateTodo)
			todos.DELETE("/:id", deleteTodo)
		}

		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Todo List API is running"})
		})
	}

	// Start server
	log.Println("Starting Todo List API server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
		sys.Exit(1)
	}
}
