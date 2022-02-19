package repository

import (
	"GoBlogClean/config"
	"GoBlogClean/models"
)

type articleRepository struct {
	config.DBHandler
}

func NewArticleRepository(dbHandler *config.DBHandler) models.ArticleRepository {
	return &articleRepository{*dbHandler}
}