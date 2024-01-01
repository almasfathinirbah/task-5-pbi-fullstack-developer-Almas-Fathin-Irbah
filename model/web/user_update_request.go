package web

type UserUpdateRequest struct {
	Id       string `validate:"required" json:"id"`
	UserName string `validate:"required, min = 1, max = 50" json:"username"`
}
