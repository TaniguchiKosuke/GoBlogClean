package repository

import (
	"GoBlogClean/article"
	"GoBlogClean/config"
	"GoBlogClean/domain"
)

type articleRepository struct {
	config.DBHandler
}

func NewArticleRepository(dbHandler *config.DBHandler) article.ArticleRepository {
	return &articleRepository{*dbHandler}
}

func (ar *articleRepository) PostArticle(article *domain.Article) (*domain.Article, error) {
	if err := ar.Conn.Create(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (ar *articleRepository) GetArticles() ([]*domain.Article, error) {
	var articles []*domain.Article
	if err := ar.Conn.Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}

func (ar *articleRepository) GetArticleByID(articleID int) (*domain.Article, error) {
	var article *domain.Article
	if err := ar.Conn.Where("id = ?", articleID).Find(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}
