package usecase

import (
	"CleanCodeGo/user/domain/entity"
	portRepository "CleanCodeGo/user/domain/port"
	"CleanCodeGo/user/usecase/exception"
	portUseCase "CleanCodeGo/user/usecase/port"
	"CleanCodeGo/user/usecase/request"
	"CleanCodeGo/user/usecase/response"
)

type CreateUser struct {
	repository portRepository.UserRepository
}

func NewCreateUser(repository portRepository.UserRepository) portUseCase.ICreateUser {
	return &CreateUser{repository: repository}
}

func (createUser *CreateUser) formatRequestToEntity(userRequest request.UserRequest) entity.User {
	return entity.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}
}

func (createUser *CreateUser) validateCreateUser(userRequest request.UserRequest) error {
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

func (createUser *CreateUser) Create(userRequest request.UserRequest) (response.UserResponse, error) {
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