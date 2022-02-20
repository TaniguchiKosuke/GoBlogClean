package handler

import "github.com/gin-gonic/gin"

func InitUserRouting(engine *gin.Engine, handler *UserHandler) {
	engine.POST("/auth/login", handler.Login)
	engine.POST("/auth/signup", handler.Signup)
	engine.GET("/auth/users", handler.GetUsers)
}