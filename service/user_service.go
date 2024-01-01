package service

import (
	"context"
	"task-5-pbi-btpns-almas/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId string)
	FindById(ctx context.Context, userId string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}