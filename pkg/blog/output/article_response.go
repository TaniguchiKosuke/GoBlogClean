package output

import (
	userOutput "GoBlogClean/pkg/auth/output"
)

type ArticleListResponse struct {
	ArticleResponse []*ArticleResponse `json:"articles"`
}

type ArticleResponse struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`

	Title    string                  `json:"title"`
	Content  string                  `json:"content"`
	AuthorID string                  `json:"user_id"`
	Author   userOutput.UserResponse `json:"user"`
	Comments []CommentResponse       `json:"comments"`
}
