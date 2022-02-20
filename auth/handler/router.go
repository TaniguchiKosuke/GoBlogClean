package handler

import "github.com/gin-gonic/gin"

func InitUserRouting(engin *gin.Engine, handler *UserHandler) {
	engin.POST("/auth/login", handler.Login)
	engin.POST("/auth/signup", handler.Signup)
}