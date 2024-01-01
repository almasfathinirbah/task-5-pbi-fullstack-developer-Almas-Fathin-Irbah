package web

type UserCreateRequest struct {
	UserName  string `validate:"required, min = 1, max = 50" json:"username"`
	Email     string `validate:"required, min = 1, max = 50, email" json:"email"`
	Password  string `validate:"required" json:"password"`
}