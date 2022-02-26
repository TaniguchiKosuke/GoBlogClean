package usecase

import (
	"GoBlogClean/pkg/blog"
	"GoBlogClean/domain"
)

type articleUsecase struct {
	articleRepository blog.ArticleRepository
}

func NewArticleUsecase(articleRepository blog.ArticleRepository) blog.ArticleUsecase {
	return &articleUsecase{articleRepository: articleRepository}
}

func (au *articleUsecase) PostArticle(article *domain.Article) error {
	_, err := au.articleRepository.PostArticle(article)
	if err != nil {
		return err
	}

	return nil
}

func (au *articleUsecase) GetArticles() ([]*domain.Article, error) {
	articles, err := au.articleRepository.GetArticles()
	if err != nil {
		return articles, err
	}

	return articles, nil
}

func (au *articleUsecase) GetArticleByID(articleID int) (*domain.Article, error) {
	article, err := au.articleRepository.GetArticleByID(articleID)
	if err != nil {
		return article, err
	}

	return article, nil
}
