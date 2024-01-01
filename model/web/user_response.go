package web

type UserResponse struct {
	Id         string `json:"id"`
	UserName   string `json:"username"`
	Email      string `json:"email"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
