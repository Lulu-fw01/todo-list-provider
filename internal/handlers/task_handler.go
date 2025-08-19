package handlers

import (
	"net/http"
	"strconv"

	"todo-list-provider/internal/auth"
	"todo-list-provider/internal/models"
	"todo-list-provider/internal/services"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	userID := auth.GetUserIDFromCtx(c)

	tasks, err := h.taskService.GetTasksByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responses []models.TaskResponse
	for _, task := range tasks {
		responses = append(responses, task.ToResponse())
	}

	c.JSON(http.StatusOK, responses)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	userID := auth.GetUserIDFromCtx(c)

	task, err := h.taskService.GetTaskByID(id, userID)
	if err != nil {
		if err == models.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found", "id": id})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task.ToResponse())
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req models.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := auth.GetUserIDFromCtx(c)

	task, err := h.taskService.CreateTask(&req, userID)
	if err != nil {
		if err == models.ErrEmptyTitle || err == models.ErrInvalidPriority {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task.ToResponse())
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var req models.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := auth.GetUserIDFromCtx(c)

	task, err := h.taskService.UpdateTask(id, &req, userID)
	if err != nil {
		if err == models.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found", "id": id})
			return
		}
		if err == models.ErrEmptyTitle || err == models.ErrInvalidPriority {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task.ToResponse())
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	userID := auth.GetUserIDFromCtx(c)

	err = h.taskService.DeleteTask(id, userID)
	if err != nil {
		if err == models.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found", "id": id})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully", "id": id})
}
