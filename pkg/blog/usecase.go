package blog

import (
	"GoBlogClean/domain"
	"GoBlogClean/pkg/blog/input"
)

type ArticleUsecase interface {
	PostArticle(article *input.ArticleRequest) error
	GetArticleByID(articleID int) (*domain.Article, error)
	GetArticles() ([]*domain.Article, error)
	// UpdateArticle(article *Article) (*Article, error)
	// DeleteArticle(article *Article) (*Article, error)
	// SearchArticle(text string) ([]*Article, error)
}
