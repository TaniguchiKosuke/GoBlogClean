package injector

import (
	"GoBlogClean/pkg/article"
	"GoBlogClean/pkg/article/handler"
	"GoBlogClean/pkg/article/repository"
	"GoBlogClean/pkg/article/usecase"
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
