package routers

import "github.com/gin-gonic/gin"

type Option func(*gin.RouterGroup)

var options []Option

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init 初始化
func Init() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api").Group("/v1")
	for _, opt := range options {
		opt(api)
	}
	return r
}
