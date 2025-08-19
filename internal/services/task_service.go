package services

import (
	"fmt"
	"time"

	"todo-list-provider/internal/models"
	"todo-list-provider/internal/repositories"
)

type TaskService struct {
	taskRepo *repositories.TaskRepository
}

func NewTaskService(taskRepo *repositories.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(req *models.CreateTaskRequest, userID int) (*models.Task, error) {
	if err := s.validateCreateTaskRequest(req); err != nil {
		return nil, err
	}

	now := time.Now()
	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		UserID:      userID,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := task.Validate(); err != nil {
		return nil, err
	}

	createdTask, err := s.taskRepo.CreateTask(task)
	if err != nil {
		return nil, fmt.Errorf("failed to create task in database: %w", err)
	}

	return createdTask, nil
}

func (s *TaskService) GetTaskByID(id int, userID int) (*models.Task, error) {
	task, err := s.taskRepo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	if task.UserID != userID {
		return nil, models.ErrTaskNotFound
	}

	return task, nil
}

func (s *TaskService) GetTasksByUserID(userID int) ([]*models.Task, error) {
	tasks, err := s.taskRepo.GetTasksByUserID(userID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) UpdateTask(id int, req *models.UpdateTaskRequest, userID int) (*models.Task, error) {
	existingTask, err := s.taskRepo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	if existingTask.UserID != userID {
		return nil, models.ErrTaskNotFound
	}

	if req.Title != nil {
		existingTask.Title = *req.Title
	}
	if req.Description != nil {
		existingTask.Description = *req.Description
	}
	if req.Completed != nil {
		existingTask.Completed = *req.Completed
	}
	if req.Priority != nil {
		existingTask.Priority = *req.Priority
	}
	if req.DueDate != nil {
		existingTask.DueDate = req.DueDate
	}

	existingTask.UpdatedAt = time.Now()

	if err := existingTask.Validate(); err != nil {
		return nil, err
	}

	if err := s.taskRepo.UpdateTask(existingTask); err != nil {
		return nil, err
	}

	return existingTask, nil
}

func (s *TaskService) DeleteTask(id int, userID int) error {
	return s.taskRepo.DeleteTask(id, userID)
}

func (s *TaskService) validateCreateTaskRequest(req *models.CreateTaskRequest) error {
	if req.Title == "" {
		return models.ErrEmptyTitle
	}

	if req.Priority < 1 || req.Priority > 5 {
		return models.ErrInvalidPriority
	}

	return nil
}
