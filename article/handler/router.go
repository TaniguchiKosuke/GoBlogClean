package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouting(engine *gin.Engine, handler *ArticleHandler) {
	engine.GET("/")
}