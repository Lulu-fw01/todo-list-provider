package main

import (
	"log"
	"net/http"
	"os"
	"todo-list-provider/configs"
	"todo-list-provider/internal/auth"
	"todo-list-provider/internal/database"
	"todo-list-provider/internal/handlers"
	"todo-list-provider/internal/repositories"
	"todo-list-provider/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	configs := configs.LoadConfig()

	db, err := database.NewConnection(&configs.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		os.Exit(1)
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db.DB)
	taskRepository := repositories.NewTaskRepository(db.DB)

	userService := services.NewUserService(userRepository)
	taskService := services.NewTaskService(taskRepository)
	authService := auth.NewAuthService(&configs.Auth)

	userHandler := handlers.NewUserHandler(&authService, userService)
	taskHandler := handlers.NewTaskHandler(taskService)

	api := r.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Todo List API is running"})
		})

		usersGr := api.Group("/user")
		usersGr.POST("/register", userHandler.RegisterUser)
		usersGr.POST("/login", userHandler.LoginUser)

		tasksGr := api.Group("/task")
		tasksGr.Use(authService.JWTAuthMiddleware())
		tasksGr.GET("/", taskHandler.GetTasks)
		tasksGr.GET("/:id", taskHandler.GetTask)
		tasksGr.POST("/", taskHandler.CreateTask)
		tasksGr.PUT("/:id", taskHandler.UpdateTask)
		tasksGr.DELETE("/:id", taskHandler.DeleteTask)

	}

	log.Println("Starting Todo List API server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
		os.Exit(1)
	}
}
