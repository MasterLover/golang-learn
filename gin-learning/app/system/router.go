package system

import (
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.RouterGroup) {
	api := r.Group("/system")

	api.GET("/ping", Ping)
	api.POST("/login", Login)
	api.DELETE("/delete", Login)
	api.PUT("/update", Login)
	api.POST("/register", Register)
}
