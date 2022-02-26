package output

type CommentResponse struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`

	Content   string `json:"content"`
	AuthorID  string `json:"user_id"`
	Author    string `json:"user"`
	ArticleID int    `json:"article_id"`
}
