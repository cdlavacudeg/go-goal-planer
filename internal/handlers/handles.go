package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cdlavacudeg/go-goal-planer/internal/entities"
)

type userResponse struct {
	Data []entities.User `json:"data"`
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users from the database
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} userResponse
// @Router /users [get]
func GetUsers(context *gin.Context) {
	users := entities.GetUsers()
	context.JSON(http.StatusOK, userResponse{Data: users})
}
