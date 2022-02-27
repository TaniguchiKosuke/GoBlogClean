package blog

import (
	"GoBlogClean/pkg/blog/input"
	"GoBlogClean/pkg/blog/output"
)

type ArticleUsecase interface {
	PostArticle(article *input.PostArticleRequest) error
	GetArticleByID(articleID int) (*output.ArticleResponse, error)
	GetArticles() (*output.GetArticlesResponse, error)
	// UpdateArticle(article *Article) (*Article, error)
	// DeleteArticle(article *Article) (*Article, error)
	// SearchArticle(text string) ([]*Article, error)
}
