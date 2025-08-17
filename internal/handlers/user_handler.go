package handlers

import (
	"net/http"
	"todo-list-provider/internal/auth"
	"todo-list-provider/internal/models"
	"todo-list-provider/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
	authService *auth.AuthServcie
}

func NewUserHandler(authService *auth.AuthServcie, userService *services.UserService) *UserHandler {
	return &UserHandler{
		authService: authService,
		userService: userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req models.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.createResponse(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := req.Email
	user, err := h.userService.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	passwordHashExpected := user.Password
	correct := auth.CheckPasswordHash(req.Password, passwordHashExpected)
	if !correct {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect password"})
		return
	}

	resp, err := h.createResponse(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) createResponse(user *models.User) (*models.RegisterUserResponse, error) {
	token, err := h.authService.CreateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.RegisterUserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Token: token,
	}, nil
}
