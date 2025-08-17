package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"` // "-" means this field won't be included in JSON responses
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type RegisterUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type RegisterUserResponse struct {
	ID    int    `json:"id" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required"`
	Token string `json:"jwt" binding:"required"`
}

type UpdateUserRequest struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Password *string `json:"password"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return ErrEmptyEmail
	}

	if u.Password == "" {
		return ErrEmptyPassword
	}

	if u.Name == "" {
		return ErrEmptyName
	}

	return nil
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
