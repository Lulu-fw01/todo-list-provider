package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"todo-list-provider/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	query := `
		INSERT INTO task (title, description, completed, priority, due_date, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`

	var id int
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.Priority,
		task.DueDate,
		task.UserID,
		task.CreatedAt,
		task.UpdatedAt,
	).Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	task.ID = id
	task.CreatedAt = createdAt
	task.UpdatedAt = updatedAt

	return task, nil
}

func (r *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	query := `
		SELECT id, title, description, completed, priority, due_date, user_id, created_at, updated_at
		FROM task
		WHERE id = $1
	`

	task := &models.Task{}
	err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.Priority,
		&task.DueDate,
		&task.UserID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrTaskNotFound
		}
		return nil, fmt.Errorf("failed to get task by ID: %w", err)
	}

	return task, nil
}

func (r *TaskRepository) GetTasksByUserID(userID int) ([]*models.Task, error) {
	query := `
		SELECT id, title, description, completed, priority, due_date, user_id, created_at, updated_at
		FROM task
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks by user ID: %w", err)
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Priority,
			&task.DueDate,
			&task.UserID,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over tasks: %w", err)
	}

	return tasks, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	query := `
		UPDATE task 
		SET title = $1, description = $2, completed = $3, priority = $4, due_date = $5, updated_at = $6
		WHERE id = $7 AND user_id = $8
	`

	result, err := r.db.Exec(
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.Priority,
		task.DueDate,
		task.UpdatedAt,
		task.ID,
		task.UserID,
	)

	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return models.ErrTaskNotFound
	}

	return nil
}

func (r *TaskRepository) DeleteTask(id int, userID int) error {
	query := `DELETE FROM task WHERE id = $1 AND user_id = $2`

	result, err := r.db.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return models.ErrTaskNotFound
	}

	return nil
}
