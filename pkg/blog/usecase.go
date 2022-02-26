package blog

import (
	"GoBlogClean/domain"
	"GoBlogClean/pkg/blog/input"
	"GoBlogClean/pkg/blog/output"
)

type ArticleUsecase interface {
	PostArticle(article *input.ArticleRequest) error
	GetArticleByID(articleID int) (*domain.Article, error)
	GetArticles() (*output.ArticleListResponse, error)
	// UpdateArticle(article *Article) (*Article, error)
	// DeleteArticle(article *Article) (*Article, error)
	// SearchArticle(text string) ([]*Article, error)
}
