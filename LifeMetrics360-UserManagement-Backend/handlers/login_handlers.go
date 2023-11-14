package handlers

import (
	"github.com/fzambone/LifeMetrics360-UserManagement/models"
	"github.com/fzambone/LifeMetrics360/LifeMetrics360-Utils-Backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Login function to authenticate user and issue JWT
func (h *Handlers) Login(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	authenticatedUser, err := h.UserService.ValidateUserCredentials(c.Request().Context(), user.Username, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateToken(authenticatedUser.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
