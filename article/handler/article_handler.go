package handler

import (
	"GoBlogClean/models"
)

type ArticleHandler struct {
	articleUsecase models.ArticleUsecase
}

func NewArticleHandler(articleUsecase models.ArticleUsecase) ArticleHandler {
	return ArticleHandler{articleUsecase: articleUsecase}
}