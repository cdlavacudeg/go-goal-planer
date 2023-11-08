package main

import (
	"github.com/gin-gonic/gin"

	"github.com/cdlavacudeg/go-goal-planner/internal/routes"
	"github.com/cdlavacudeg/go-goal-planner/utils/env"
)

func main() {
	env.LoadConfig()
	r := gin.Default()
	r = routes.RouterApi(r)
	r.Run(":8080")
}
