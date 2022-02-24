package injector

import (
	"GoBlogClean/article"
	"GoBlogClean/article/handler"
	"GoBlogClean/article/repository"
	"GoBlogClean/article/usecase"
)

func InjectArticleRepository() article.ArticleRepository {
	return repository.NewArticleRepository(InjectDB())
}

func InjectArticleUsecase() article.ArticleUsecase {
	return usecase.NewArticleUsecase(InjectArticleRepository())
}

func InjectArticleHandler() handler.ArticleHandler {
	return handler.NewArticleHandler(InjectArticleUsecase())
}
