package handler

import (
	"GoBlogClean/pkg/blog"
	"GoBlogClean/pkg/blog/input"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleUsecase blog.ArticleUsecase
}

func NewArticleHandler(articleUsecase blog.ArticleUsecase) ArticleHandler {
	return ArticleHandler{articleUsecase: articleUsecase}
}

func (ah *ArticleHandler) PostArticle(c *gin.Context) {
	var requestBody *input.ArticleRequest
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ah.articleUsecase.PostArticle(requestBody); err != nil {
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
