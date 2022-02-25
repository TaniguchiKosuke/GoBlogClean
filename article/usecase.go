package article

import "GoBlogClean/domain"

type ArticleUsecase interface {
	PostArticle(article *domain.Article) error
	GetArticleByID(articleID int) (*domain.Article, error)
	GetArticles() ([]*domain.Article, error)
	// UpdateArticle(article *Article) (*Article, error)
	// DeleteArticle(article *Article) (*Article, error)
	// SearchArticle(text string) ([]*Article, error)
}
