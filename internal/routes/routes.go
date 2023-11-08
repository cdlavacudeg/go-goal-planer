package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/cdlavacudeg/go-goal-planner/docs"
)

func RouterApi(router *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"
	// @title Goal planner api
	// @version 1.0
	// @description This is a rest api for goal plannig
	// @BasePath /api/v1

	// @securityDefinitions.apikey apiKeyAuth
	// @in header
	// @name Authorization
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		usersRouter(v1)
	}

	return router
}
