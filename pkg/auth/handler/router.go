package handler

import (
	"GoBlogClean/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouting(engine *gin.Engine, handler *UserHandler) {
	engine.POST("/auth/login", handler.Login)
	engine.POST("/auth/signup", handler.Signup)

	loginRequired := engine.Group("/", middleware.Authenticate())
	{
		loginRequired.GET("/auth/users", handler.GetUsers)
		loginRequired.PUT("auth/update", handler.UpdateUsername)
	}
}
