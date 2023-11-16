package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cdlavacudeg/go-goal-planner/internal/entities"
)

type userResponse struct {
	Data []entities.User `json:"data"`
}

type errorResponse struct {
	Error string `json:"error"`
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users from the database
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} userResponse
// @Router /users [get]
// func GetUsers(context *gin.Context) {
// 	users := entities.GetUsers()
// 	context.JSON(http.StatusOK, userResponse{Data: users})
// }

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Security apiKeyAuth
// @Param user body entities.User true "User"
// @Success 201 {object} entities.User
// @Failure 400 {object} errorResponse
// @Router /users [post]
func CreateUser(context *gin.Context) {
	var req entities.User

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	service := GetService()
	err := entities.CreateItem(*service, req, "GoalPlanner")
	if err != nil {
		context.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	context.JSON(http.StatusCreated, req)
}
