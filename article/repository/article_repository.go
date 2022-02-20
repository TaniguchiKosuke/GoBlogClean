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

func (ar *articleRepository) PostArticle(article *models.Article) (*models.Article, error) {
	if err := ar.Conn.Create(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}