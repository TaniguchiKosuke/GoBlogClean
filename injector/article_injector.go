package injector

import (
	"GoBlogClean/article/handler"
	"GoBlogClean/article/repository"
	"GoBlogClean/article/usecase"
	"GoBlogClean/models"
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
