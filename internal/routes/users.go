package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/cdlavacudeg/go-goal-planner/internal/handlers"
)

func usersRouter(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/users")
	{
		// users.GET("", handlers.GetUsers)

		// users.POST godoc
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
		users.POST("", handlers.CreateUser)
	}
}
