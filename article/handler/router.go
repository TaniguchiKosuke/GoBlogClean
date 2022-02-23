package handler

import (
	"github.com/gin-gonic/gin"
)

func InitArticleRouting(engine *gin.Engine, handler *ArticleHandler) {
	engine.POST("/article/create", handler.PostArticle)
	engine.GET("/", handler.GetArticles)
	engine.GET("/article/:id", handler.GetArticleByID)
}
