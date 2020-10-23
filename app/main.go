package main

import (
	"CleanCodeGo/infrastructure/createuser/repository/inmemory"
	"CleanCodeGo/interfaceadapters/createuser/controller"
	"CleanCodeGo/interfaceadapters/createuser/controller/request"
	"CleanCodeGo/usecases/createuser"
	"fmt"
)

func main() {

	//remoteRepository := mysql.NewUserRepository(&gorm.DB{})
	inMemoryRepository := inmemory.NewUserInMemoryRepository()

	useCase := createuser.NewCreateUser(inMemoryRepository)
	newUserController := controller.NewUserController(useCase)

	userRequest := request.UserRequest{Name: "David", Email: "David@", Password: "123"}

	response, err := newUserController.CreateUser(userRequest)

	if err != nil {
		panic(err)
	}

	fmt.Sprint(response)
}
