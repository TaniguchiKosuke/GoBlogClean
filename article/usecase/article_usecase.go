package usecase

import "GoBlogClean/models"

type articleUsecase struct {
	articleRepository models.ArticleRepository
}

func NewArticleUsecase(articleRepository models.ArticleRepository) models.ArticleUsecase {
	return &articleUsecase{articleRepository: articleRepository}
}