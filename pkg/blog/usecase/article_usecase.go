package usecase

import (
	"GoBlogClean/domain"
	userOutput "GoBlogClean/pkg/auth/output"
	"GoBlogClean/pkg/blog"
	"GoBlogClean/pkg/blog/input"
	blogOutput "GoBlogClean/pkg/blog/output"
	"GoBlogClean/pkg/constant"
)

type articleUsecase struct {
	articleRepository blog.ArticleRepository
}

func NewArticleUsecase(articleRepository blog.ArticleRepository) blog.ArticleUsecase {
	return &articleUsecase{articleRepository: articleRepository}
}

func (au *articleUsecase) PostArticle(requestBody *input.ArticleRequest) error {
	// authorなどもここでdomainに渡す
	article := &domain.Article{
		Title:   requestBody.Title,
		Content: requestBody.Content,
	}

	_, err := au.articleRepository.PostArticle(article)
	if err != nil {
		return err
	}

	return nil
}

func (au *articleUsecase) GetArticles() (*blogOutput.ArticleListResponse, error) {
	articles, err := au.articleRepository.GetArticles()
	if err != nil {
		return &blogOutput.ArticleListResponse{}, nil
	}

	articleList := make([]*blogOutput.ArticleResponse, 0, len(articles))

	for _, article := range articles {
		author := &userOutput.UserResponse{
			ID:        article.AuthorID,
			CreatedAt: article.Author.CreatedAt.Format(constant.TimeLayout),
			UpdatedAt: article.Author.UpdatedAt.Format(constant.TimeLayout),
			DeletedAt: article.Author.DeletedAt.Time.Format(constant.TimeLayout),
			Username:  article.Author.Username,
		}

		articleResponse := &blogOutput.ArticleResponse{
			ID:        article.ID,
			CreatedAt: article.CreatedAt.Format(constant.TimeLayout),
			UpdatedAt: article.UpdatedAt.Format(constant.TimeLayout),
			DeletedAt: article.DeletedAt.Time.Format(constant.TimeLayout),
			Title:     article.Title,
			Content:   article.Content,
			AuthorID:  article.AuthorID,
			Author:    *author,
			// Comments: ,
		}

		articleList = append(articleList, articleResponse)
	}

	articleListResponse := &blogOutput.ArticleListResponse{
		ArticleResponse: articleList,
	}

	return articleListResponse, nil
}

func (au *articleUsecase) GetArticleByID(articleID int) (*blogOutput.ArticleResponse, error) {
	article, err := au.articleRepository.GetArticleByID(articleID)
	if err != nil {
		return &blogOutput.ArticleResponse{}, err
	}

	author := &userOutput.UserResponse{
		ID:        article.Author.ID,
		CreatedAt: article.Author.CreatedAt.Format(constant.TimeLayout),
		UpdatedAt: article.Author.UpdatedAt.Format(constant.TimeLayout),
		DeletedAt: article.Author.DeletedAt.Time.Format(constant.TimeLayout),
		Username:  article.Author.Username,
	}

	articleResponse := &blogOutput.ArticleResponse{
		ID:        article.ID,
		CreatedAt: article.CreatedAt.Format(constant.TimeLayout),
		UpdatedAt: article.UpdatedAt.Format(constant.TimeLayout),
		DeletedAt: article.DeletedAt.Time.Format(constant.TimeLayout),
		Title:     article.Title,
		Content:   article.Content,
		AuthorID:  article.AuthorID,
		Author:    *author,
		// Comments: ,
	}

	return articleResponse, nil
}
