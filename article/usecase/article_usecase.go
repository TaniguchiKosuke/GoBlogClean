package usecase

import (
	"GoBlogClean/article"
	"GoBlogClean/models"
)

type articleUsecase struct {
	articleRepository article.ArticleRepository
}

func NewArticleUsecase(articleRepository article.ArticleRepository) article.ArticleUsecase {
	return &articleUsecase{articleRepository: articleRepository}
}

func (au *articleUsecase) PostArticle(article *models.Article) error {
	_, err := au.articleRepository.PostArticle(article)
	if err != nil {
		return err
	}

	return nil
}

func (au *articleUsecase) GetArticles() ([]*models.Article, error) {
	articles, err := au.articleRepository.GetArticles()
	if err != nil {
		return articles, err
	}

	return articles, nil
}

func (au *articleUsecase) GetArticleByID(articleID int) (*models.Article, error) {
	article, err := au.articleRepository.GetArticleByID(articleID)
	if err != nil {
		return article, err
	}

	return article, nil
}
