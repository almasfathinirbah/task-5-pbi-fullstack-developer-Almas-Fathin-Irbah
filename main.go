package main

import (
	"net/http"
	"task-5-pbi-btpns-almas/controller"
	"task-5-pbi-btpns-almas/database"
	"task-5-pbi-btpns-almas/helper"
	"task-5-pbi-btpns-almas/repository"
	"task-5-pbi-btpns-almas/service"
	
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	// connection to database
	db := database.Database()

	// validation
	validate := validator.New()

	//repository
	userRepository := repository.NewUserRepositoryImpl()

	//service
	userService := service.NewUserServiceImpl(
		userRepository,
		db,
		validate,
	)

	// controller
	userController := controller.NewUserControllerImpl(userService)

	// initialize
	router := httprouter.New()

	// router
	// [USER]
	router.POST("/api/v1/user", userController.Create)
	router.PUT("/api/v1/user", userController.Update)
	router.DELETE("/api/v1/user/:user_id", userController.Delete)
	router.GET("/api/v1/user/:user_id", userController.FindById)
	router.GET("/api/v1/user", userController.FindAll)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}