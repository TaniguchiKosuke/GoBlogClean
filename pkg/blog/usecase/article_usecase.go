package usecase

import (
	"time"

	"GoBlogClean/domain"
	userOutput "GoBlogClean/pkg/auth/output"
	"GoBlogClean/pkg/blog"
	"GoBlogClean/pkg/blog/input"
	blogOutput "GoBlogClean/pkg/blog/output"
)

type articleUsecase struct {
	articleRepository blog.ArticleRepository
}

func NewArticleUsecase(articleRepository blog.ArticleRepository) blog.ArticleUsecase {
	return &articleUsecase{articleRepository: articleRepository}
}

func (au *articleUsecase) PostArticle(articleRequest *input.PostArticleRequest) error {
	article := &domain.Article{
		Title:    articleRequest.Title,
		Content:  articleRequest.Content,
		AuthorID: articleRequest.AuthorID,
	}

	_, err := au.articleRepository.PostArticle(article)
	if err != nil {
		return err
	}

	return nil
}

func (au *articleUsecase) GetArticles() (*blogOutput.GetArticlesResponse, error) {
	articles, err := au.articleRepository.GetArticles()
	if err != nil {
		return &blogOutput.GetArticlesResponse{}, nil
	}

	articleList := make([]*blogOutput.ArticleResponse, 0, len(articles))

	for _, article := range articles {
		author := &userOutput.UserResponse{
			ID:        article.AuthorID,
			CreatedAt: article.Author.CreatedAt.Format(time.RFC3339),
			UpdatedAt: article.Author.UpdatedAt.Format(time.RFC3339),
			DeletedAt: article.Author.DeletedAt.Time.Format(time.RFC3339),
			Username:  article.Author.Username,
		}

		articleResponse := &blogOutput.ArticleResponse{
			ID:        article.ID,
			CreatedAt: article.CreatedAt.Format(time.RFC3339),
			UpdatedAt: article.UpdatedAt.Format(time.RFC3339),
			DeletedAt: article.DeletedAt.Time.Format(time.RFC3339),
			Title:     article.Title,
			Content:   article.Content,
			AuthorID:  article.AuthorID,
			Author:    *author,
			// Comments: ,
		}

		articleList = append(articleList, articleResponse)
	}

	articleListResponse := &blogOutput.GetArticlesResponse{
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
		CreatedAt: article.Author.CreatedAt.Format(time.RFC3339),
		UpdatedAt: article.Author.UpdatedAt.Format(time.RFC3339),
		DeletedAt: article.Author.DeletedAt.Time.Format(time.RFC3339),
		Username:  article.Author.Username,
	}

	articleResponse := &blogOutput.ArticleResponse{
		ID:        article.ID,
		CreatedAt: article.CreatedAt.Format(time.RFC3339),
		UpdatedAt: article.UpdatedAt.Format(time.RFC3339),
		DeletedAt: article.DeletedAt.Time.Format(time.RFC3339),
		Title:     article.Title,
		Content:   article.Content,
		AuthorID:  article.AuthorID,
		Author:    *author,
		// Comments: ,
	}

	return articleResponse, nil
}
