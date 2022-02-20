package handler

import (
	"GoBlogClean/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleUsecase models.ArticleUsecase
}

func NewArticleHandler(articleUsecase models.ArticleUsecase) ArticleHandler {
	return ArticleHandler{articleUsecase: articleUsecase}
}

func (ah *ArticleHandler) PostArticle(c *gin.Context) {
	var article *models.Article
	if err := c.BindJSON(&article); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, article)
		return
	}

	if err := ah.articleUsecase.PostArticle(article); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, article)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "successed"})
}

func (ah *ArticleHandler) GetArticles(c *gin.Context) {
	articles, err := ah.articleUsecase.GetArticles()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, articles)
		return
	}

	c.JSON(http.StatusOK, articles)
}