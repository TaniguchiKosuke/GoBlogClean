package injector

import (
	"GoBlogClean/pkg/blog"
	"GoBlogClean/pkg/blog/handler"
	"GoBlogClean/pkg/blog/repository"
	"GoBlogClean/pkg/blog/usecase"
)

func InjectArticleRepository() blog.ArticleRepository {
	return repository.NewArticleRepository(InjectDB())
}

func InjectArticleUsecase() blog.ArticleUsecase {
	return usecase.NewArticleUsecase(InjectArticleRepository())
}

func InjectArticleHandler() handler.ArticleHandler {
	return handler.NewArticleHandler(InjectArticleUsecase())
}
