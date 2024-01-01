package utils

import (
	"task-5-pbi-btpns-almas/model/entity"
	"task-5-pbi-btpns-almas/model/web"
)

func UserResponse(user entity.Users) web.UserResponse {
	return web.UserResponse{
		Id: user.Id,
		UserName: user.UserName,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}