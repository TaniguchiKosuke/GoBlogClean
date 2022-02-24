package article

import "GoBlogClean/models"

type ArticleRepository interface {
	PostArticle(article *models.Article) (*models.Article, error)
	GetArticleByID(articleID int) (*models.Article, error)
	GetArticles() ([]*models.Article, error)
	// UpdateArticle(article *Article) (*Article, error)
	// DeleteArticle(article *Article) (*Article, error)
	// SearchArticle(text string) ([]*Article, error)
}
