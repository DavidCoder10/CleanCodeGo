package inmemory

import (
	"CleanCodeGo/entities/user"
	"CleanCodeGo/usecases/createuser/gateway"
	"errors"
)

type UserRepositoryInMemory struct {
	users map[string]user.User
}

func NewUserInMemoryRepository() gateway.Gateway {
	return &UserRepositoryInMemory{users: map[string]user.User{}}
}

func (userRepository *UserRepositoryInMemory) FindByEmail(email string) (user.User, error) {
	if foundUser, ok := userRepository.users[email]; ok {
		return foundUser, nil
	}
	return user.User{}, errors.New("user not found")
}

func (userRepository *UserRepositoryInMemory) Create(user user.User) (user.User, error) {
	userRepository.users[user.Email] = user
	return user, nil
}
