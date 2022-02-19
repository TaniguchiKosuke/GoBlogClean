package injector

import (
	"GoBlogClean/models"
	"GoBlogClean/article/usecase"
	"GoBlogClean/article/handler"
	"GoBlogClean/article/repository"
)

func InjectArticleRepository() models.ArticleRepository {
	return repository.NewArticleRepository(InjectDB())
}

func InjectArticleUsecase() models.ArticleUsecase {
	return usecase.NewArticleUsecase(InjectArticleRepository())
}

func InjectArticleHandler() handler.ArticleHandler {
	return handler.NewArticleHandler(InjectArticleUsecase())
}