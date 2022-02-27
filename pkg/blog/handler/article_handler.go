package handler

import (
	"GoBlogClean/pkg/blog"
	"GoBlogClean/pkg/blog/input"
	"GoBlogClean/pkg/util"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleUsecase blog.ArticleUsecase
}

func NewArticleHandler(articleUsecase blog.ArticleUsecase) ArticleHandler {
	return ArticleHandler{articleUsecase: articleUsecase}
}

func (ah *ArticleHandler) PostArticle(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		log.Println("Authorization is empty")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Authorization is empty"})
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := util.GetUserIDFromJWT(tokenStr)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, "token is not valid")
		return
	}
	if userID == "" {
		log.Println("failed to get user id")
		c.JSON(http.StatusBadRequest, "failed to get user id")
		return
	}

	var requestBody *input.PostArticleRequest
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	requestBody.AuthorID = userID

	if err := ah.articleUsecase.PostArticle(requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": " succeeded"})
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
