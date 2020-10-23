package controller

import (
	"CleanCodeGo/interfaceadapters/createuser/controller/request"
	"CleanCodeGo/usecases/createuser"
	userRequestUseCase "CleanCodeGo/usecases/createuser/request"
)

type UserController struct {
	useCaseUser createuser.Interactor
}

func NewUserController(useCaseUser createuser.Interactor) *UserController {
	return &UserController{useCaseUser: useCaseUser}
}

func (userController *UserController) CreateUser(request request.UserRequest) (string, error) {
	userRequestUseCase := userRequestUseCase.UserRequest{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	_, err := userController.useCaseUser.Create(userRequestUseCase)
	if err != nil {
		return "Error", err
	}
	return "Success", err
}
