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

func (ar *articleRepository) GetArticles() ([]*models.Article, error) {
	var articles []*models.Article
	if err := ar.Conn.Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}

func (ar *articleRepository) GetArticleByID(articleID int) (*models.Article, error) {
	var article *models.Article
	if err := ar.Conn.Where("id = ?", articleID).Find(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}
