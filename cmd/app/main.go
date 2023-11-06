package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	docs "github.com/cdlavacudeg/go-goal-planner/docs"
	"github.com/cdlavacudeg/go-goal-planner/internal/handlers"
	"github.com/cdlavacudeg/go-goal-planner/utils/env"
)

// @title Goal planner api
// @version 1.0
// @description This is a rest api for goal plannig
// @BasePath /api/v1

// @securityDefinitions.apikey apiKeyAuth
// @in header
// @name Authorization
func main() {
	env.LoadConfig()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
		users := v1.Group("/users")
		{
			users.GET("", handlers.GetUsers)
			users.POST("", handlers.CreateUser)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

// Helloworld godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "Hello World!")
}
