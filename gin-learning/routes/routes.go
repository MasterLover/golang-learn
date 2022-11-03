package routes

import (
	user "gin-learning/user/router"

	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine) {
	api := engine.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/ping", user.Ping)
	v1.POST("/login", user.Login)
	v1.DELETE("/delete", user.Login)
	v1.PUT("/update", user.Login)
	v1.POST("/register", user.Register)
}
