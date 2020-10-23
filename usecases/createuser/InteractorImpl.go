package createuser

import (
	"CleanCodeGo/entities/user"
	"CleanCodeGo/usecases/createuser/exception"
	"CleanCodeGo/usecases/createuser/gateway"
	"CleanCodeGo/usecases/createuser/request"
	"CleanCodeGo/usecases/createuser/response"
)

type InteractorImpl struct {
	repository gateway.Gateway
}

func NewCreateUser(repository gateway.Gateway) Interactor {
	return &InteractorImpl{repository: repository}
}

func (createUser *InteractorImpl) formatRequestToEntity(userRequest request.UserRequest) user.User {
	return user.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}
}

func (createUser *InteractorImpl) validateCreateUser(userRequest request.UserRequest) error {
	var userAlreadyExistsException *exception.UserAlreadyExistsException
	_, err := createUser.repository.FindByEmail(userRequest.Email)
	if err != nil {
		return err
	}
	err = userAlreadyExistsException.UserAlreadyExistsException("email exists")
	if err != nil {
		return err
	}
	return err
}

func (createUser *InteractorImpl) Create(userRequest request.UserRequest) (response.UserResponse, error) {
	var userResponse response.UserResponse
	err := createUser.validateCreateUser(userRequest)
	if err != nil {
		return userResponse, err
	}
	userEntity := createUser.formatRequestToEntity(userRequest)
	user, err := createUser.repository.Create(userEntity)
	if err != nil {
		return userResponse, err
	}
	userResponse.Name = user.Name
	return userResponse, err
}
