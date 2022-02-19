package handler

import "github.com/gin-gonic/gin"

func InitUserRouting(engin *gin.Engine, handler *UserHandler) {
	engin.GET("auth/")
}