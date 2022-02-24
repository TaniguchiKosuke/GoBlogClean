package article

import "GoBlogClean/models"

type ArticleUsecase interface {
	PostArticle(article *models.Article) error
	GetArticleByID(articleID int) (*models.Article, error)
	GetArticles() ([]*models.Article, error)
	// UpdateArticle(article *Article) (*Article, error)
	// DeleteArticle(article *Article) (*Article, error)
	// SearchArticle(text string) ([]*Article, error)
}
