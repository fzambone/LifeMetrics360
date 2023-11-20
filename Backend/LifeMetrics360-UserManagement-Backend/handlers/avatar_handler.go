package handlers

import (
	"github.com/fzambone/LifeMetrics360-UserManagement/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AvatarHandlers struct {
	AvatarService *services.AvatarService
}

func NewAvatarHandlers(avatarService *services.AvatarService) *AvatarHandlers {
	return &AvatarHandlers{AvatarService: avatarService}
}

func (h *AvatarHandlers) Avatar(c echo.Context) error {
	userEmail := c.Param("email")

	// Search for user by email and get the avatar URL
	avatarURL, err := h.AvatarService.GetAvatarByEmail(c.Request().Context(), userEmail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"avatarURL": avatarURL})
}
