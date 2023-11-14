package handlers

import (
	"github.com/fzambone/LifeMetrics360-UserManagement/models"
	"github.com/fzambone/LifeMetrics360-UserManagement/services"
	"github.com/fzambone/LifeMetrics360/LifeMetrics360-Utils-Backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handlers struct {
	UserService *services.UserService
}

func NewHandlers(userService *services.UserService) *Handlers {
	return &Handlers{UserService: userService}
}

func (h *Handlers) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// TODO: Implement validation logic

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error hashing password")
	}
	user.Password = hashedPassword

	// Save the user to the database
	userID, err := h.UserService.CreateUser(c.Request().Context(), *user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user.ID = userID
	return c.JSON(http.StatusCreated, user)
}

func (h *Handlers) GetUser(c echo.Context) error {
	userID := c.Param("id")

	// Retrieving user from database
	user, err := h.UserService.GetUserByID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handlers) UpdateUser(c echo.Context) error {
	userID := c.Param("id")
	updateData := new(models.User)
	if err := c.Bind(updateData); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Update user in database
	err := h.UserService.UpdateUser(c.Request().Context(), userID, updateData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User updated successfully")
}

func (h *Handlers) DeleteUser(c echo.Context) error {
	userID := c.Param("id")

	// Delete user from database
	if err := h.UserService.DeleteUser(c.Request().Context(), userID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User deleted successfully")
}
