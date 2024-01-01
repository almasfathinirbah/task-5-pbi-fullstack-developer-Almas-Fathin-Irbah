package service

import (
	"context"
	"database/sql"
	"task-5-pbi-btpns-almas/helper"
	"task-5-pbi-btpns-almas/model/entity"
	"task-5-pbi-btpns-almas/model/web"
	"task-5-pbi-btpns-almas/repository"
	"task-5-pbi-btpns-almas/utils"
	"time"

	"github.com/go-playground/validator"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	helper.Defer(tx)

	passwordHash, err := utils.HashPassword(request.Password)
	helper.PanicError(err)

	user := entity.Users{
		Id:        utils.Uuid(),
		UserName:  request.UserName,
		Email:     request.Email,
		Password:  passwordHash,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	user = service.UserRepository.Create(
		ctx,
		tx,
		user,
	)

	return utils.UserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	helper.Defer(tx)

	user, err := service.UserRepository.FindById(
		ctx,
		tx,
		request.Id,
	)
	helper.PanicError(err)

	user.UserName = request.UserName
	user.UpdatedAt = time.Now().Unix()

	user = service.UserRepository.Update(
		ctx,
		tx,
		user,
	)

	return utils.UserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId string) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	helper.Defer(tx)

	user, err := service.UserRepository.FindById(
		ctx,
		tx,
		userId,
	)
	helper.PanicError(err)

	service.UserRepository.Delete(
		ctx,
		tx,
		user,
	)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	helper.Defer(tx)

	user, err := service.UserRepository.FindById(
		ctx,
		tx,
		userId,
	)
	helper.PanicError(err)

	return utils.UserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	helper.Defer(tx)

	users := service.UserRepository.FindAll(
		ctx,
		tx,
	)

	var UserResponses []web.UserResponse
	for _, user := range users {
		UserResponses = append(UserResponses, utils.UserResponse(user))
	}

	return UserResponses
}
