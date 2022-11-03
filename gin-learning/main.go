package main

import (
	"gin-learning/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	routes.InitRoutes(engine)
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
