package handler

import (
	"GoBlogClean/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitArticleRouting(engine *gin.Engine, handler *ArticleHandler) {
	loginRequired := engine.Group("/", middleware.Authenticate())
	{
		loginRequired.POST("/article/create", handler.PostArticle)
		loginRequired.GET("/articles", handler.GetArticles)
		loginRequired.GET("/article/:id", handler.GetArticleByID)
	}
}
