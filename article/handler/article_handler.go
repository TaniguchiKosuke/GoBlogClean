package handler

import (
	"GoBlogClean/models"
	"log"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ah.articleUsecase.PostArticle(article); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": " succeeded"})
}

func (ah *ArticleHandler) GetArticles(c *gin.Context) {
	articles, err := ah.articleUsecase.GetArticles()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, articles)
}

func (ah *ArticleHandler) GetArticleByID(c *gin.Context) {
	articleIDString := c.Param("id")
	articleID, err := strconv.Atoi(articleIDString)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
	}

	article, err := ah.articleUsecase.GetArticleByID(articleID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, article)
}
