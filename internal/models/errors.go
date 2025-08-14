package models

import "errors"

var (
	ErrEmptyTitle      = errors.New("todo title cannot be empty")
	ErrInvalidPriority = errors.New("priority must be between 1 and 5")
	ErrTodoNotFound    = errors.New("todo not found")
	ErrInvalidUserID   = errors.New("invalid user ID")
	ErrInvalidTodoID   = errors.New("invalid todo ID")
	ErrUnauthorized    = errors.New("unauthorized access")
)

var (
	ErrEmptyEmail       = errors.New("email cannot be empty")
	ErrEmptyPassword    = errors.New("password cannot be empty")
	ErrTooShortPassword = errors.New("password is too short")
	ErrEmptyName        = errors.New("name cannot be empty")
	ErrUserNotFound     = errors.New("user not found")
	ErrUserExists       = errors.New("user already exists")
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrInvalidPassword  = errors.New("invalid password")
)
