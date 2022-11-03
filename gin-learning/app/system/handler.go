package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Ping(ctx *gin.Context) {
	_, err := ctx.Writer.Write([]byte("pong"))
	if err != nil {
		return
	}
}

type LoginForm struct {
	Username string `form:"username" bind:"required"`
	Password string `form:"password" bind:"required"`
}

func Login(ctx *gin.Context) {
	var loginForm LoginForm
	err := ctx.ShouldBind(&loginForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	time.Sleep(1 * time.Second)
	if loginForm.Username != "root" || loginForm.Password != "admin" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"STATUS": 401,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
	})

}

func Register(ctx *gin.Context) {

}
