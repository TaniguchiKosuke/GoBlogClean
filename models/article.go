package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Title    string    `json:"title"`
	Content  string    `json:"content"`
	AuthorID string    `json:"user_id"`
	Author   User      `json:"user"`
	Comments []Comment `json:"comment" gorm:"foreignKey:AuthorID"`
}

type ArticleRepository interface {
	PostArticle(article *Article) (*Article, error)
	GetArticleByID(articleID int) (*Article, error)
	GetArticles() ([]*Article, error)
	// UpdateArticle(article *Article) (*Article, error)
	// DeleteArticle(article *Article) (*Article, error)
	// SearchArticle(text string) ([]*Article, error)
}

type ArticleUsecase interface {
	PostArticle(article *Article) error
	GetArticleByID(articleID int) (*Article, error)
	GetArticles() ([]*Article, error)
	// UpdateArticle(article *Article) (*Article, error)
	// DeleteArticle(article *Article) (*Article, error)
	// SearchArticle(text string) ([]*Article, error)
}
