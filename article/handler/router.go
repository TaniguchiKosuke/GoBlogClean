package handler

import (
	"github.com/gin-gonic/gin"
)

func InitArticleRouting(engine *gin.Engine, handler *ArticleHandler) {
	engine.GET("/")
}