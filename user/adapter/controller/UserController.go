package controller

import (
	"CleanCodeGo/user/adapter/controller/request"
	"CleanCodeGo/user/usecase/port"
	userRequestUseCase "CleanCodeGo/user/usecase/request"
)

type UserController struct {
	useCaseUser port.ICreateUser
}

func NewUserController(useCaseUser port.ICreateUser) *UserController {
	return &UserController{useCaseUser: useCaseUser}
}

func (userController *UserController) CreateUser(request request.UserRequest) (string, error) {
	userRequestUseCase := userRequestUseCase.UserRequest{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}
	_, err := userController.useCaseUser.Create(userRequestUseCase)
	if err != nil {
		return "Error", err
	}
	return "Success", err
}