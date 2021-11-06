package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controllers"
	"web_app/logger"
	"web_app/middleware"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)

	// 登录业务路由
	v1.POST("/login", controllers.LoginHandler)

	v1.Use(middleware.JWTAuthMiddleware())

	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
	}

	// 验证路由
	r.GET("/validate", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		// 判断请求头中是否有 有效的JWT
		c.String(http.StatusOK, "验证成功")
	})
	return r
}
