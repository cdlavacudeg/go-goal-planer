package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/cdlavacudeg/go-goal-planner/internal/handlers"
)

func usersRouter(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/users")
	{
		users.GET("", handlers.GetUsers)
		users.POST("", handlers.CreateUser)
	}
}
