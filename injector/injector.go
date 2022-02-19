package injector

import (
	"GoBlogClean/config"
	"GoBlogClean/models"
	"GoBlogClean/article/usecase"
	"GoBlogClean/article/handler"
	"GoBlogClean/article/repository"
)

func InjectDB() *config.DBHandler {
	dbHandler := config.NewDBHandler()
	return dbHandler
}

func InjectArticleRepository() models.ArticleRepository {
	return repository.NewArticleRepository(InjectDB())
}

func InjectArticleUsecase() models.ArticleUsecase {
	return usecase.NewArticleUsecase(InjectArticleRepository())
}

func InjectArticleHandler() handler.ArticleHandler {
	return handler.NewArticleHandler(InjectArticleUsecase())
}