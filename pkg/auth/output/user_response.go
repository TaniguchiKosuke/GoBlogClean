package output

type UserResponse struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`

	Username string `json:"username"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetUsersResponse struct {
	Users []*UserResponse `json:"users"`
}
