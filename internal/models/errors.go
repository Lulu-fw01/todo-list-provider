package models

import "errors"

var (
	ErrEmptyTitle      = errors.New("task title cannot be empty")
	ErrInvalidPriority = errors.New("priority must be between 1 and 5")
	ErrTaskNotFound    = errors.New("task not found")
	ErrInvalidUserID   = errors.New("invalid user ID")
	ErrInvalidTaskID   = errors.New("invalid task ID")
	ErrUnauthorized    = errors.New("unauthorized access")
)

var (
	ErrEmptyEmail       = errors.New("email cannot be empty")
	ErrEmptyPassword    = errors.New("password cannot be empty")
	ErrNotValidPassword = errors.New("not valid password")
	ErrEmptyName        = errors.New("name cannot be empty")
	ErrUserNotFound     = errors.New("user not found")
	ErrUserExists       = errors.New("user already exists")
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrInvalidPassword  = errors.New("invalid password")
)
